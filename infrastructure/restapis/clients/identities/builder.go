package identities

import (
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	"github.com/xmn-services/buckets-network/application/commands/identities/chains"
	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	"github.com/xmn-services/buckets-network/infrastructure/restapis/shared"
)

type builder struct {
	bucketBuilder  buckets.Builder
	storageBuilder storages.Builder
	chainBuilder   chains.Builder
	minerBuilder   miners.Builder
	client         *resty.Client
	peer           peer.Peer
	name           string
	password       string
	seed           string
}

func createBuilder(
	bucketBuilder buckets.Builder,
	storageBuilder storages.Builder,
	chainBuilder chains.Builder,
	minerBuilder miners.Builder,
	client *resty.Client,
	peer peer.Peer,
) identities.Builder {
	out := builder{
		bucketBuilder:  bucketBuilder,
		storageBuilder: storageBuilder,
		chainBuilder:   chainBuilder,
		minerBuilder:   minerBuilder,
		client:         client,
		peer:           peer,
		name:           "",
		password:       "",
		seed:           "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() identities.Builder {
	return createBuilder(
		app.bucketBuilder,
		app.storageBuilder,
		app.chainBuilder,
		app.minerBuilder,
		app.client,
		app.peer,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) identities.Builder {
	app.name = name
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) identities.Builder {
	app.password = password
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) identities.Builder {
	app.seed = seed
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (identities.Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build an Application instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Application instance")
	}

	bucket, err := app.bucketBuilder.Create().WithName(app.name).WithPassword(app.password).WithSeed(app.seed).Now()
	if err != nil {
		return nil, err
	}

	storage, err := app.storageBuilder.Create().WithName(app.name).WithPassword(app.password).WithSeed(app.seed).Now()
	if err != nil {
		return nil, err
	}

	chain, err := app.chainBuilder.Create().WithName(app.name).WithPassword(app.password).WithSeed(app.seed).Now()
	if err != nil {
		return nil, err
	}

	miner, err := app.minerBuilder.Create().WithName(app.name).WithPassword(app.password).WithSeed(app.seed).Now()
	if err != nil {
		return nil, err
	}

	token := shared.AuthenticateToBase64(&shared.Authenticate{
		Name:     app.name,
		Password: app.password,
		Seed:     app.seed,
	})

	subApplications := createSubApplications(bucket, storage, chain, miner)
	current := createCurrent(app.client, token, app.peer)
	return createApplication(current, subApplications), nil
}
