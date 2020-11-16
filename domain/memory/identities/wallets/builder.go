package wallets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages"
)

type builder struct {
	minerFactory   miners.Factory
	storageFactory storages.Factory
	miner          miners.Miner
	storage        storages.Storage
}

func createBuilder(
	minerFactory miners.Factory,
	storageFactory storages.Factory,
) Builder {
	out := builder{
		minerFactory:   minerFactory,
		storageFactory: storageFactory,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.minerFactory, app.storageFactory)
}

// WithMiner adds a miner to the builder
func (app *builder) WithMiner(miner miners.Miner) Builder {
	app.miner = miner
	return app
}

// WithStorage adds a storage to the builder
func (app *builder) WithStorage(storage storages.Storage) Builder {
	app.storage = storage
	return app
}

// Now builds a new Wallet instance
func (app *builder) Now() (Wallet, error) {
	if app.miner == nil {
		miner, err := app.minerFactory.Create()
		if err != nil {
			return nil, err
		}

		app.miner = miner
	}

	if app.storage == nil {
		storage, err := app.storageFactory.Create()
		if err != nil {
			return nil, err
		}

		app.storage = storage
	}

	return createWallet(app.miner, app.storage), nil
}
