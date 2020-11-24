package syncs

import (
	"github.com/xmn-services/buckets-network/application/commands"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/peers"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	clientBuilder commands.ClientBuilder,
	appli commands.Application,
	chainService chains.Service,
) Builder {
	chainBuilder := chains.NewBuilder()
	peersBuilder := peers.NewBuilder()
	return createBuilder(
		appli,
		clientBuilder,
		chainBuilder,
		chainService,
		peersBuilder,
	)
}

// Builder represents a sync application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	WithAdditionalBucketsPerBlock(additionalBuckets uint) Builder
	Now() (Application, error)
}

// Application represents the sync application
type Application interface {
	Chain() error
	Peers() error
	Storage() error
}
