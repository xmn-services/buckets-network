package identities

import (
	"errors"

	"github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
)

type builder struct {
	bucketBuilder      buckets.Builder
	storageBuilder     storages.Builder
	identityRepository identities.Repository
	identityService    identities.Service
	name               string
	password           string
	seed               string
}

func createBuilder(
	bucketBuilder buckets.Builder,
	storageBuilder storages.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
) Builder {
	out := builder{
		bucketBuilder:      bucketBuilder,
		storageBuilder:     storageBuilder,
		identityRepository: identityRepository,
		identityService:    identityService,
		name:               "",
		password:           "",
		seed:               "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.bucketBuilder,
		app.storageBuilder,
		app.identityRepository,
		app.identityService,
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

	bucketApp, err := app.bucketBuilder.Create().WithName(app.name).WithPassword(app.password).WithSeed(app.seed).Now()
	if err != nil {
		return nil, err
	}

	storageApp, err := app.storageBuilder.Create().WithName(app.name).WithPassword(app.password).WithSeed(app.seed).Now()
	if err != nil {
		return nil, err
	}

	subApps := createSubApplications(bucketApp, storageApp)
	currentApp := createCurrent(app.identityRepository, app.identityService, app.name, app.password, app.seed)
	return createApplication(currentApp, subApps), nil
}
