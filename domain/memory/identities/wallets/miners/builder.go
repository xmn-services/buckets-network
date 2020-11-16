package miners

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/permanents"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter             hash.Adapter
	mutableBuilder          entities.MutableBuilder
	blocksFactory           blocks.Factory
	bucketsFactory          buckets.Factory
	permanentBucketsFactory permanents.Factory
	hash                    *hash.Hash
	withoutHash             bool
	toTransact              buckets.Buckets
	queue                   buckets.Buckets
	broadcasted             permanents.Buckets
	toLink                  blocks.Blocks
	createdOn               *time.Time
	lastUpdatedOn           *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	mutableBuilder entities.MutableBuilder,
	blocksFactory blocks.Factory,
	bucketsFactory buckets.Factory,
	permanentBucketsFactory permanents.Factory,
) Builder {
	out := builder{
		hashAdapter:             hashAdapter,
		mutableBuilder:          mutableBuilder,
		blocksFactory:           blocksFactory,
		bucketsFactory:          bucketsFactory,
		permanentBucketsFactory: permanentBucketsFactory,
		hash:                    nil,
		withoutHash:             false,
		toTransact:              nil,
		queue:                   nil,
		broadcasted:             nil,
		toLink:                  nil,
		createdOn:               nil,
		lastUpdatedOn:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.mutableBuilder,
		app.blocksFactory,
		app.bucketsFactory,
		app.permanentBucketsFactory,
	)
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

// WithToTransact adds a toTransact Buckets to the builder
func (app *builder) WithToTransact(toTransact buckets.Buckets) Builder {
	app.toTransact = toTransact
	return app
}

// WithQueue adds a queue Buckets to the builder
func (app *builder) WithQueue(queue buckets.Buckets) Builder {
	app.queue = queue
	return app
}

// WithBroadcasted adds a broadcasted Buckets to the builder
func (app *builder) WithBroadcasted(broadcasted permanents.Buckets) Builder {
	app.broadcasted = broadcasted
	return app
}

// WithToLink adds a toLink Blocks to the builder
func (app *builder) WithToLink(toLink blocks.Blocks) Builder {
	app.toLink = toLink
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

// Now builds a new Miner instance
func (app *builder) Now() (Miner, error) {
	if app.toTransact == nil {
		toTransact, err := app.bucketsFactory.Create()
		if err != nil {
			return nil, err
		}

		app.toTransact = toTransact
	}

	if app.queue == nil {
		queue, err := app.bucketsFactory.Create()
		if err != nil {
			return nil, err
		}

		app.queue = queue
	}

	if app.broadcasted == nil {
		broadcasted, err := app.permanentBucketsFactory.Create()
		if err != nil {
			return nil, err
		}

		app.broadcasted = broadcasted
	}

	if app.toLink == nil {
		toLink, err := app.blocksFactory.Create()
		if err != nil {
			return nil, err
		}

		app.toLink = toLink
	}

	if app.withoutHash {
		hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
			app.toTransact.Hash().Bytes(),
			app.queue.Hash().Bytes(),
			app.broadcasted.Hash().Bytes(),
			app.toLink.Hash().Bytes(),
		})

		if err != nil {
			return nil, err
		}

		app.hash = hsh
	}

	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Miner instance")
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	return createMiner(mutable, app.toTransact, app.queue, app.broadcasted, app.toLink), nil
}
