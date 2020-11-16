package miners

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/permanents"
)

type builder struct {
	blocksFactory           blocks.Factory
	bucketsFactory          buckets.Factory
	permanentBucketsFactory permanents.Factory
	toTransact              buckets.Buckets
	queue                   buckets.Buckets
	broadcasted             permanents.Buckets
	toLink                  blocks.Blocks
}

func createBuilder(
	blocksFactory blocks.Factory,
	bucketsFactory buckets.Factory,
	permanentBucketsFactory permanents.Factory,
) Builder {
	out := builder{
		blocksFactory:           blocksFactory,
		bucketsFactory:          bucketsFactory,
		permanentBucketsFactory: permanentBucketsFactory,
		toTransact:              nil,
		queue:                   nil,
		broadcasted:             nil,
		toLink:                  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.blocksFactory,
		app.bucketsFactory,
		app.permanentBucketsFactory,
	)
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

	return createMiner(app.toTransact, app.queue, app.broadcasted, app.toLink), nil
}
