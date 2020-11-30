package storages

import (
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	"github.com/xmn-services/buckets-network/infrastructure/restapis/shared"
)

type builder struct {
	client   *resty.Client
	peer     peer.Peer
	name     string
	password string
	seed     string
}

func createBuilder(
	client *resty.Client,
	peer peer.Peer,
) storages.Builder {
	out := builder{
		client:   client,
		peer:     peer,
		name:     "",
		password: "",
		seed:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() storages.Builder {
	return createBuilder(
		app.client,
		app.peer,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) storages.Builder {
	app.name = name
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) storages.Builder {
	app.password = password
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) storages.Builder {
	app.seed = seed
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (storages.Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build an Application instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Application instance")
	}

	token, err := shared.AuthenticateToBase64(&shared.Authenticate{
		Name:     app.name,
		Password: app.password,
		Seed:     app.seed,
	})

	if err != nil {
		return nil, err
	}

	return createApplication(app.client, token, app.peer), nil
}
