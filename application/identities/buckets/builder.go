package buckets

import (
	"errors"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	identity_buckets "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter           hash.Adapter
	pkFactory             encryption.Factory
	chunkBuilder          chunks.Builder
	fileBuilder           files.Builder
	bucketBuilder         buckets.Builder
	bucketRepository      buckets.Repository
	identityBucketBuilder identity_buckets.Builder
	identityRepository    identities.Repository
	identityService       identities.Service
	chunkSizeInBytes      uint
	name                  string
	password              string
	seed                  string
}

func createBuilder(
	hashAdapter hash.Adapter,
	pkFactory encryption.Factory,
	chunkBuilder chunks.Builder,
	fileBuilder files.Builder,
	bucketBuilder buckets.Builder,
	bucketRepository buckets.Repository,
	identityBucketBuilder identity_buckets.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
	chunkSizeInBytes uint,
) Builder {
	out := builder{
		hashAdapter:           hashAdapter,
		pkFactory:             pkFactory,
		chunkBuilder:          chunkBuilder,
		fileBuilder:           fileBuilder,
		bucketBuilder:         bucketBuilder,
		bucketRepository:      bucketRepository,
		identityBucketBuilder: identityBucketBuilder,
		identityRepository:    identityRepository,
		identityService:       identityService,
		chunkSizeInBytes:      chunkSizeInBytes,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.pkFactory,
		app.chunkBuilder,
		app.fileBuilder,
		app.bucketBuilder,
		app.bucketRepository,
		app.identityBucketBuilder,
		app.identityRepository,
		app.identityService,
		app.chunkSizeInBytes,
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
		app.pkFactory,
		app.chunkBuilder,
		app.fileBuilder,
		app.bucketBuilder,
		app.bucketRepository,
		app.identityBucketBuilder,
		app.identityRepository,
		app.identityService,
		app.name,
		app.password,
		app.seed,
		app.chunkSizeInBytes,
	), nil
}
