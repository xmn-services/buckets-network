package clients

import (
	"errors"
	"fmt"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands"
	command_identities "github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	"github.com/xmn-services/buckets-network/infrastructure/restapis/shared"
)

type current struct {
	commandIdentityBuilder command_identities.Builder
	client                 *resty.Client
	url                    string
}

func createCurrent(
	commandIdentityBuilder command_identities.Builder,
	client *resty.Client,
	peer peer.Peer,
) commands.Current {
	out := current{
		commandIdentityBuilder: commandIdentityBuilder,
		client:                 client,
		url:                    fmt.Sprintf("%s%s", peer.String(), "/identities"),
	}

	return &out
}

// NewIdentity creates a new identity instance
func (app *current) NewIdentity(name string, password string, seed string, root string) error {
	resp, err := app.client.R().
		SetFormDataFromValues(shared.IdentityToURLValues(shared.Identity{
			Authenticate: &shared.Authenticate{
				Name:     name,
				Password: password,
				Seed:     seed,
			},
			Root: root,
		})).
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
	cmdIdentityApp, err := app.commandIdentityBuilder.Create().
		WithName(name).
		WithPassword(password).
		WithSeed(seed).
		Now()

	if err != nil {
		return nil, err
	}

	_, err = cmdIdentityApp.Current().Retrieve()
	if err != nil {
		return nil, err
	}

	return cmdIdentityApp, nil
}

// UpdateIdentity updates an identity instance
func (app *current) UpdateIdentity(identity identities.Identity, password string, newPassword string) error {
	auth := shared.Authenticate{
		Name:     identity.Name(),
		Password: password,
		Seed:     identity.Seed(),
	}

	token, err := shared.AuthenticateToBase64(&auth)
	if err != nil {
		return err
	}

	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, token).
		SetFormDataFromValues(shared.IdentityToURLValues(shared.Identity{
			Authenticate: &auth,
			Root:         identity.Root(),
		})).
		Put(app.url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}
