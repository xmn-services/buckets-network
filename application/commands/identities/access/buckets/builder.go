package buckets

import (
	"errors"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter        hash.Adapter
	identityRepository identities.Repository
	identityService    identities.Service
	bucketRepository   buckets.Repository
	contentService     contents.Service
	name               string
	password           string
	seed               string
	bucketHashStr      string
}

func createBuilder(
	hashAdapter hash.Adapter,
	identityRepository identities.Repository,
	identityService identities.Service,
	bucketRepository buckets.Repository,
	contentService contents.Service,
) Builder {
	out := builder{
		hashAdapter:        hashAdapter,
		identityRepository: identityRepository,
		identityService:    identityService,
		bucketRepository:   bucketRepository,
		contentService:     contentService,
		name:               "",
		password:           "",
		seed:               "",
		bucketHashStr:      "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.identityRepository,
		app.identityService,
		app.bucketRepository,
		app.contentService,
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

// WithBucket adds a bucketHash to the builder
func (app *builder) WithBucket(bucketHashStr string) Builder {
	app.bucketHashStr = bucketHashStr
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

	if app.bucketHashStr == "" {
		return nil, errors.New("the bucketHash is mandatory in order to build an Application instance")
	}

	bucketHash, err := app.hashAdapter.FromString(app.bucketHashStr)
	if err != nil {
		return nil, err
	}

	return createApplication(
		app.identityRepository,
		app.identityService,
		app.bucketRepository,
		app.contentService,
		app.name,
		app.password,
		app.seed,
		*bucketHash,
	), nil
}
