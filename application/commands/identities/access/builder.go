package access

import (
	"errors"

	"github.com/xmn-services/buckets-network/application/commands/identities/access/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses/access"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter        hash.Adapter
	accessBuilder      access.Builder
	identityRepository identities.Repository
	identityService    identities.Service
	bucketAppBuilder   buckets.Builder
	name               string
	password           string
	seed               string
}

func createBuilder(
	hashAdapter hash.Adapter,
	accessBuilder access.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
	bucketAppBuilder buckets.Builder,
) Builder {
	out := builder{
		hashAdapter:        hashAdapter,
		accessBuilder:      accessBuilder,
		identityRepository: identityRepository,
		identityService:    identityService,
		bucketAppBuilder:   bucketAppBuilder,
		name:               "",
		password:           "",
		seed:               "",
	}

	return &out
}

// Create creates a new builder instance
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.accessBuilder,
		app.identityRepository,
		app.identityService,
		app.bucketAppBuilder,
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

	return createApplication(
		app.hashAdapter,
		app.accessBuilder,
		app.identityRepository,
		app.identityService,
		app.bucketAppBuilder,
		app.name,
		app.password,
		app.seed,
	), nil
}
