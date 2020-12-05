package identities

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/access"
	"github.com/xmn-services/buckets-network/application/commands/identities/chains"
	"github.com/xmn-services/buckets-network/application/commands/identities/lists"
	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
)

type subApplications struct {
	access  access.Application
	lists   lists.Application
	storage storages.Application
	chain   chains.Application
	miner   miners.Application
}

func createSubApplications(
	access access.Application,
	lists lists.Application,
	storage storages.Application,
	chain chains.Application,
	miner miners.Application,
) SubApplications {
	out := subApplications{
		access:  access,
		lists:   lists,
		storage: storage,
		chain:   chain,
		miner:   miner,
	}

	return &out
}

// Access returns the access application
func (app *subApplications) Access() access.Application {
	return app.access
}

// List returns the list application
func (app *subApplications) List() lists.Application {
	return app.lists
}

// Storage returns the storage application
func (app *subApplications) Storage() storages.Application {
	return app.storage
}

// Chain returns the chain application
func (app *subApplications) Chain() chains.Application {
	return app.chain
}

// Miner returns the miner application
func (app *subApplications) Miner() miners.Application {
	return app.miner
}
