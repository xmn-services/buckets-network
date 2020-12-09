package identities

import (
	"github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/application/commands/identities/access"
	"github.com/xmn-services/buckets-network/application/commands/identities/chains"
	"github.com/xmn-services/buckets-network/application/commands/identities/lists"
	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
)

type subApplications struct {
	storage storages.Application
	chain   chains.Application
	miner   miners.Application
}

func createSubApplications(
	storage storages.Application,
	chain chains.Application,
	miner miners.Application,
) identities.SubApplications {
	out := subApplications{
		storage: storage,
		chain:   chain,
		miner:   miner,
	}

	return &out
}

// Access returns the access application
func (obj *subApplications) Access() access.Application {
	return nil
}

// List returns the list application
func (obj *subApplications) List() lists.Application {
	return nil
}

// Storage returns the storage application
func (obj *subApplications) Storage() storages.Application {
	return obj.storage
}

// Chain returns the chain application
func (obj *subApplications) Chain() chains.Application {
	return obj.chain
}

// Miner returns the miner application
func (obj *subApplications) Miner() miners.Application {
	return obj.miner
}
