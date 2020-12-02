package contacts

import (
	"errors"

	application_contact_bucket "github.com/xmn-services/buckets-network/application/commands/identities/lists/contacts/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter        hash.Adapter
	identityRepository identities.Repository
	identityService    identities.Service
	bucketAppBuilder   application_contact_bucket.Builder
	name               string
	password           string
	seed               string
	listHashStr        string
	contactHashStr     string
}

func createBuilder(
	hashAdapter hash.Adapter,
	identityRepository identities.Repository,
	identityService identities.Service,
	bucketAppBuilder application_contact_bucket.Builder,
) Builder {
	out := builder{
		hashAdapter:        hashAdapter,
		identityRepository: identityRepository,
		identityService:    identityService,
		bucketAppBuilder:   bucketAppBuilder,
		name:               "",
		password:           "",
		seed:               "",
		listHashStr:        "",
		contactHashStr:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
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

// WithList adds a list to the builder
func (app *builder) WithList(listHashStr string) Builder {
	app.listHashStr = listHashStr
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

	if app.listHashStr == "" {
		return nil, errors.New("the listHash is mandatory in order to build an Application instance")
	}

	listHash, err := app.hashAdapter.FromString(app.listHashStr)
	if err != nil {
		return nil, err
	}

	return createApplication(
		app.hashAdapter,
		app.identityRepository,
		app.identityService,
		app.bucketAppBuilder,
		app.name,
		app.password,
		app.seed,
		*listHash,
	), nil
}
