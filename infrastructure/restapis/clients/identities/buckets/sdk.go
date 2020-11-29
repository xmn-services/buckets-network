package buckets

import (
	"github.com/go-resty/resty/v2"
	command_bucket "github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewBuilder creates a new builder instance
func NewBuilder(peer peer.Peer) command_bucket.Builder {
	bucketAdapter := buckets.NewAdapter()
	client := resty.New()
	return createBuilder(bucketAdapter, client, peer)
}
