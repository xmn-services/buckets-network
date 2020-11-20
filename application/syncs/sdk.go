package syncs

import (
	"github.com/xmn-services/buckets-network/application/commands"
	application_chain "github.com/xmn-services/buckets-network/application/commands/chains"
	application_identity_buckets "github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	application_identity_storages "github.com/xmn-services/buckets-network/application/commands/identities/storages"
	application_storages "github.com/xmn-services/buckets-network/application/commands/identities/storages"
	application_miners "github.com/xmn-services/buckets-network/application/commands/miners"
	application_peers "github.com/xmn-services/buckets-network/application/commands/peers"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/peers"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	chainApp application_chain.Application,
	minerApp application_miners.Application,
	peersApp application_peers.Application,
	storageApp application_storages.Application,
	identityBucketApp application_identity_buckets.Application,
	identityStorageApp application_identity_storages.Application,
	identityRepository identities.Repository,
	identityService identities.Service,
	chainService chains.Service,
	clientBuilder commands.ClientBuilder,
) Builder {
	chainBuilder := chains.NewBuilder()
	peersBuilder := peers.NewBuilder()
	return createBuilder(
		chainApp,
		minerApp,
		peersApp,
		storageApp,
		identityBucketApp,
		identityStorageApp,
		identityRepository,
		identityService,
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
	Miners() error
	MinerBlock() error
	MinerLink() error
}
