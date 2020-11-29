package identities

import (
	"github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	"github.com/xmn-services/buckets-network/application/commands/identities/chains"
	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewBuilder creates a new builder instace
func NewBuilder(
	bucketBuilder buckets.Builder,
	storageBuilder storages.Builder,
	chainBuilder chains.Builder,
	minerBuilder miners.Builder,
	peer peer.Peer,
) identities.Builder {
	client := resty.New()
	return createBuilder(bucketBuilder, storageBuilder, chainBuilder, minerBuilder, client, peer)
}
