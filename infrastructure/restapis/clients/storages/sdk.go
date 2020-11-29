package storages

import (
	"github.com/go-resty/resty/v2"
	commands_storages "github.com/xmn-services/buckets-network/application/commands/storages"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewApplication creates a new application instance
func NewApplication(peer peer.Peer) commands_storages.Application {
	client := resty.New()
	return createApplication(client, peer)
}
