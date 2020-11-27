package identities

import (
	"github.com/xmn-services/buckets-network/application/commands/identities"
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
) identities.SubApplications {
	out := subApplications{
		bucket:  bucket,
		storage: storage,
		chain:   chain,
		miner:   miner,
	}

	return &out
}

// Bucket returns the bucket application
func (obj *subApplications) Bucket() buckets.Application {
	return obj.bucket
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
