package syncs

import (
	"errors"

	"github.com/xmn-services/buckets-network/application/commands"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/peers"
)

type builder struct {
	appli                     commands.Application
	clientBuilder             commands.ClientBuilder
	chainBuilder              chains.Builder
	chainService              chains.Service
	peersBuilder              peers.Builder
	name                      string
	password                  string
	seed                      string
	additionalBucketsPerBlock uint
}

func createBuilder(
	appli commands.Application,
	clientBuilder commands.ClientBuilder,
	chainBuilder chains.Builder,
	chainService chains.Service,
	peersBuilder peers.Builder,
) Builder {
	out := builder{
		appli:                     appli,
		clientBuilder:             clientBuilder,
		chainBuilder:              chainBuilder,
		chainService:              chainService,
		peersBuilder:              peersBuilder,
		name:                      "",
		password:                  "",
		seed:                      "",
		additionalBucketsPerBlock: uint(0),
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.appli,
		app.clientBuilder,
		app.chainBuilder,
		app.chainService,
		app.peersBuilder,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) Builder {
	app.password = password
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// WithAdditionalBucketsPerBlock adds an additional buckets per block to the builder
func (app *builder) WithAdditionalBucketsPerBlock(additionalBucketsPerBlock uint) Builder {
	app.additionalBucketsPerBlock = additionalBucketsPerBlock
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build an Application instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.appli,
		app.clientBuilder,
		app.chainBuilder,
		app.chainService,
		app.peersBuilder,
		app.name,
		app.password,
		app.seed,
		app.additionalBucketsPerBlock,
	), nil
}
