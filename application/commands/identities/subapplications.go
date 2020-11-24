package identities

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	"github.com/xmn-services/buckets-network/application/commands/identities/chains"
	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
)

type subApplications struct {
	bucket  buckets.Application
	storage storages.Application
	chain   chains.Application
	miner   miners.Application
}

func createSubApplications(
	bucket buckets.Application,
	storage storages.Application,
	chain chains.Application,
	miner miners.Application,
) SubApplications {
	out := subApplications{
		bucket:  bucket,
		storage: storage,
		chain:   chain,
		miner:   miner,
	}

	return &out
}

// Bucket returns the bucket application
func (app *subApplications) Bucket() buckets.Application {
	return app.bucket
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
