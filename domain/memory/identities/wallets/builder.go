package wallets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages"
)

type builder struct {
	minerFactory    miners.Factory
	storageFactory  storages.Factory
	accessesFactory accesses.Factory
	listsFactory    lists.Factory
	miner           miners.Miner
	storage         storages.Storage
	accesses        accesses.Accesses
	lists           lists.Lists
}

func createBuilder(
	minerFactory miners.Factory,
	storageFactory storages.Factory,
	accessesFactory accesses.Factory,
	listsFactory lists.Factory,
) Builder {
	out := builder{
		minerFactory:    minerFactory,
		storageFactory:  storageFactory,
		accessesFactory: accessesFactory,
		listsFactory:    listsFactory,
		miner:           nil,
		storage:         nil,
		accesses:        nil,
		lists:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.minerFactory,
		app.storageFactory,
		app.accessesFactory,
		app.listsFactory,
	)
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

// WithAccesses adds an accesses to the builder
func (app *builder) WithAccesses(accesses accesses.Accesses) Builder {
	app.accesses = accesses
	return app
}

// WithLists adds a lists to the builder
func (app *builder) WithLists(lists lists.Lists) Builder {
	app.lists = lists
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

	if app.accesses == nil {
		accesses, err := app.accessesFactory.Create()
		if err != nil {
			return nil, err
		}

		app.accesses = accesses
	}

	if app.lists == nil {
		lists, err := app.listsFactory.Create()
		if err != nil {
			return nil, err
		}

		app.lists = lists
	}

	return createWallet(app.miner, app.storage, app.accesses, app.lists), nil
}
