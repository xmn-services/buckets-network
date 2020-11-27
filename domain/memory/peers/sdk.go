package peers

import (
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	"github.com/xmn-services/buckets-network/libs/file"
)

// NewService creates a new service
func NewService(
	fileService file.Service,
	fileNameWithExt string,
) Service {
	return createService(fileService, fileNameWithExt)
}

// NewRepository creates a new repository
func NewRepository(
	fileRepository file.Repository,
	fileNameWithExt string,
) Repository {
	return createRepository(fileRepository, fileNameWithExt)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a peers adapter
type Adapter interface {
	JSONToPeers(js []byte) (Peers, error)
}

// Builder represents peers builder
type Builder interface {
	Create() Builder
	WithPeers(peers []peer.Peer) Builder
	Now() (Peers, error)
}

// Peers represents peers
type Peers interface {
	All() []peer.Peer
}

// Repository represents a peers repository
type Repository interface {
	Exists() bool
	Retrieve() (Peers, error)
}

// Service represents a peer service
type Service interface {
	Save(peers Peers) error
}
