package contents

import (
	"errors"

	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	bucket *hash.Hash
	file   *hash.Hash
	chunk  *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		bucket: nil,
		file:   nil,
		chunk:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithBucket adds a bucket to the builder
func (app *builder) WithBucket(bucket hash.Hash) Builder {
	app.bucket = &bucket
	return app
}

// WithFile adds a file to the builder
func (app *builder) WithFile(file hash.Hash) Builder {
	app.file = &file
	return app
}

// WithChunk adds a chunk to the builder
func (app *builder) WithChunk(chunk hash.Hash) Builder {
	app.chunk = &chunk
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.bucket == nil {
		return nil, errors.New("the bucket hash is mandatory in order to build a Content instance")
	}

	if app.file == nil {
		return nil, errors.New("the file hash is mandatory in order to build a Content instance")
	}

	if app.chunk == nil {
		return nil, errors.New("the chunk hash is mandatory in order to build a Content instance")
	}

	return createContent(*app.bucket, *app.file, *app.chunk), nil
}
