package peers

import (
	"github.com/xmn-services/buckets-network/domain/memory/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewApplication creates a new application
func NewApplication(
	peersRepository peers.Repository,
	peersService peers.Service,
) Application {
	peerBuilder := peer.NewBuilder()
	peersBuilder := peers.NewBuilder()
	return createApplication(
		peerBuilder,
		peersBuilder,
		peersRepository,
		peersService,
	)
}

// Application represents the peer application
type Application interface {
	Retrieve() (peers.Peers, error)
	SaveClear(host string, port uint) error
	SaveOnion(host string, port uint) error
	Save(peers peers.Peers) error
}
