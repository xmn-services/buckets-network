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
		SetBody(shared.Identity{
			Authenticate: &shared.Authenticate{
				Name:     name,
				Password: password,
				Seed:     seed,
			},
			Root: root,
		}).
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
	token := shared.AuthenticateToBase64(&shared.Authenticate{
		Name:     name,
		Password: password,
		Seed:     seed,
	})

	resp, err := app.client.R().
		SetHeader("X-Session-Token", token).
		SetResult(&shared.Authenticate{}).
		Get(app.url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		auth := resp.Result().(*shared.Authenticate)
		return app.commandIdentityBuilder.Create().
			WithName(auth.Name).
			WithPassword(auth.Password).
			WithSeed(auth.Seed).
			Now()
	}

	return nil, errors.New(string(resp.Body()))
}

// UpdateIdentity updates an identity instance
func (app *current) UpdateIdentity(identity identities.Identity, password string, newPassword string) error {
	auth := shared.Authenticate{
		Name:     identity.Name(),
		Password: password,
		Seed:     identity.Seed(),
	}

	token := shared.AuthenticateToBase64(&auth)
	update := shared.Identity{
		Authenticate: &auth,
		Root:         identity.Root(),
	}

	resp, err := app.client.R().
		SetHeader("X-Session-Token", token).
		SetBody(&update).
		Put(app.url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}
