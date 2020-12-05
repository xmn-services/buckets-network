package lists

import (
	"errors"

	"github.com/xmn-services/buckets-network/application/commands/identities/lists/contacts"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter        hash.Adapter
	listBuilder        list.Builder
	identityRepository identities.Repository
	identityService    identities.Service
	contactsAppBuilder contacts.Builder
	name               string
	password           string
	seed               string
}

func createBuilder(
	hashAdapter hash.Adapter,
	listBuilder list.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
	contactsAppBuilder contacts.Builder,
) Builder {
	out := builder{
		hashAdapter:        hashAdapter,
		listBuilder:        listBuilder,
		identityRepository: identityRepository,
		identityService:    identityService,
		contactsAppBuilder: contactsAppBuilder,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.listBuilder,
		app.identityRepository,
		app.identityService,
		app.contactsAppBuilder,
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
		app.listBuilder,
		app.identityRepository,
		app.identityService,
		app.contactsAppBuilder,
		app.name,
		app.password,
		app.seed,
	), nil
}
