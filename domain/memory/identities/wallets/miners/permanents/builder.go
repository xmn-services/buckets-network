package permanents

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	mutableBuilder entities.MutableBuilder
	hash           *hash.Hash
	withoutHash    bool
	lst            []bucket.Bucket
	createdOn      *time.Time
	lastUpdatedOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	mutableBuilder entities.MutableBuilder,
) Builder {
	out := builder{
		hashAdapter:    hashAdapter,
		mutableBuilder: mutableBuilder,
		hash:           nil,
		withoutHash:    false,
		lst:            nil,
		createdOn:      nil,
		lastUpdatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.mutableBuilder)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
}

// WithoutHash flags the builder as without hash
func (app *builder) WithoutHash() Builder {
	app.withoutHash = true
	return app
}

// WithBuckets add buckets to the builder
func (app *builder) WithBuckets(buckets []bucket.Bucket) Builder {
	app.lst = buckets
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// LastUpdatedOn adds a lastUpdatedOn time to the builder
func (app *builder) LastUpdatedOn(lastUpdatedOn time.Time) Builder {
	app.lastUpdatedOn = &lastUpdatedOn
	return app
}

// Now builds a new buckets instance
func (app *builder) Now() (Buckets, error) {
	if app.lst == nil {
		app.lst = []bucket.Bucket{}
	}

	if app.withoutHash {
		blocks := [][]byte{
			[]byte("initial block"),
		}

		for _, oneBlock := range app.lst {
			blocks = append(blocks, oneBlock.Hash().Bytes())
		}

		hsh, err := app.hashAdapter.FromMultiBytes(blocks)
		if err != nil {
			return nil, err
		}

		app.hash = hsh
	}

	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Buckets instance")
	}

	mp := map[string]bucket.Bucket{}
	for _, oneBucket := range app.lst {
		keyname := oneBucket.Hash().String()
		mp[keyname] = oneBucket
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	return crateBuckets(mutable, app.lst, mp), nil
}
