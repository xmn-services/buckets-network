package clients

import (
	"errors"
	"fmt"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands"
	command_identities "github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/application/servers/authenticates"
	server_identities "github.com/xmn-services/buckets-network/application/servers/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

type current struct {
	commandIdentityBuilder command_identities.Builder
	updateIdentityAdapter  command_identities.UpdateAdapter
	updateBuilder          command_identities.UpdateBuilder
	serverIdentityBuilder  server_identities.Builder
	authenticateBuilder    authenticates.Builder
	authenticateAdapter    authenticates.Adapter
	identityAdapter        identities.Adapter
	client                 *resty.Client
	url                    string
}

func createCurrent(
	commandIdentityBuilder command_identities.Builder,
	updateIdentityAdapter command_identities.UpdateAdapter,
	updateBuilder command_identities.UpdateBuilder,
	serverIdentityBuilder server_identities.Builder,
	authenticateBuilder authenticates.Builder,
	authenticateAdapter authenticates.Adapter,
	identityAdapter identities.Adapter,
	client *resty.Client,
	peer peer.Peer,
) commands.Current {
	out := current{
		commandIdentityBuilder: commandIdentityBuilder,
		updateIdentityAdapter:  updateIdentityAdapter,
		updateBuilder:          updateBuilder,
		serverIdentityBuilder:  serverIdentityBuilder,
		authenticateBuilder:    authenticateBuilder,
		authenticateAdapter:    authenticateAdapter,
		identityAdapter:        identityAdapter,
		client:                 client,
		url:                    fmt.Sprintf("%s%s", peer.String(), "/identities"),
	}

	return &out
}

// NewIdentity creates a new identity instance
func (app *current) NewIdentity(name string, password string, seed string, root string) error {
	auth, err := app.authenticateBuilder.Create().WithName(name).WithPassword(password).WithSeed(seed).Now()
	if err != nil {
		return err
	}

	identity, err := app.serverIdentityBuilder.Create().WithAuthenticate(auth).WithRoot(root).Now()
	if err != nil {
		return err
	}

	resp, err := app.client.R().
		SetBody(identity).
		Post(app.url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}

// Authenticate authenticates on the identity application
func (app *current) Authenticate(name string, seed string, password string) (command_identities.Application, error) {
	auth, err := app.authenticateBuilder.Create().WithName(name).WithPassword(password).WithSeed(seed).Now()
	if err != nil {
		return nil, err
	}

	token, err := app.authenticateAdapter.AuthenticateToBase64(auth)
	if err != nil {
		return nil, err
	}

	resp, err := app.client.R().
		SetHeader("X-Session-Token", token).
		Get(app.url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		js := resp.Body()
		identity, err := app.identityAdapter.JSONToIdentity(js)
		if err != nil {
			return nil, err
		}

		retName := identity.Name()
		retSeed := identity.Seed()
		return app.commandIdentityBuilder.Create().WithName(retName).WithPassword(password).WithSeed(retSeed).Now()
	}

	return nil, errors.New(string(resp.Body()))
}

// UpdateIdentity updates an identity instance
func (app *current) UpdateIdentity(identity identities.Identity, password string, newPassword string) error {
	name := identity.Name()
	seed := identity.Seed()
	auth, err := app.authenticateBuilder.Create().WithName(name).WithPassword(password).WithSeed(seed).Now()
	if err != nil {
		return err
	}

	token, err := app.authenticateAdapter.AuthenticateToBase64(auth)
	if err != nil {
		return err
	}

	update, err := app.updateBuilder.Create().WithPassword(newPassword).Now()
	if err != nil {
		return err
	}

	urlValues := app.updateIdentityAdapter.UpdateToURLValues(update)
	resp, err := app.client.R().
		SetHeader("X-Session-Token", token).
		SetBody(urlValues).
		Put(app.url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}
