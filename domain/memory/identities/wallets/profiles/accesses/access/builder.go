package access

import (
	"errors"

	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	bucket *hash.Hash
	key    public.Key
}

func createBuilder() Builder {
	out := builder{
		bucket: nil,
		key:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithBucket adds a bucket hash to the builder
func (app *builder) WithBucket(bucket hash.Hash) Builder {
	app.bucket = &bucket
	return app
}

// WithKey adds a public key to the builder
func (app *builder) WithKey(key public.Key) Builder {
	app.key = key
	return app
}

// Now builds a new Access instance
func (app *builder) Now() (Access, error) {
	if app.bucket == nil {
		return nil, errors.New("the bucket hash is mandatory in order to build an Access instance")
	}

	if app.key == nil {
		return nil, errors.New("the PublicKey is mandatory in order to build an Access instance")
	}

	return createAccess(*app.bucket, app.key), nil
}
