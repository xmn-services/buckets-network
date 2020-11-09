package transactions

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/transactions/addresses"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	address          addresses.Address
	bucket           *hash.Hash
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		address:          nil,
		bucket:           nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithAddress adds an address to the builder
func (app *builder) WithAddress(address addresses.Address) Builder {
	app.address = address
	return app
}

// WithBucket adds a bucket to the builder
func (app *builder) WithBucket(bucket hash.Hash) Builder {
	app.bucket = &bucket
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {

	if app.bucket == nil {
		return nil, errors.New("the bucket hash is mandatory in order to build a Transaction instance")
	}

	data := [][]byte{
		app.bucket.Bytes(),
	}

	if app.address != nil {
		data = append(data, app.address.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.address != nil {
		return createTransactionWithAddress(immutable, *app.bucket, app.address), nil
	}

	return createTransaction(immutable, *app.bucket), nil
}
