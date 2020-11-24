package chains

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	minerApplication miners.Application,
	identityRepository identities.Repository,
	identityService identities.Service,
	genesisRepository genesis.Repository,
	genesisService genesis.Service,
	blockService blocks.Service,
	linkService links.Service,
	chainRepository chains.Repository,
	chainService chains.Service,
) Builder {
	genesisBuilder := genesis.NewBuilder()
	blockBuilder := blocks.NewBuilder()
	minedBlockBuilder := mined_block.NewBuilder()
	linkBuilder := links.NewBuilder()
	minedLinkBuilder := mined_link.NewBuilder()
	chainBuilder := chains.NewBuilder()
	return createBuilder(
		minerApplication,
		identityRepository,
		identityService,
		genesisBuilder,
		genesisRepository,
		genesisService,
		blockBuilder,
		blockService,
		minedBlockBuilder,
		linkBuilder,
		linkService,
		minedLinkBuilder,
		chainBuilder,
		chainRepository,
		chainService,
	)
}

// Builder represents a chain application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// Application represents the chain application
type Application interface {
	Init(
		miningValue uint8,
		baseDifficulty uint,
		increasePerBucket float64,
		linkDifficulty uint,
		rootAdditionalBuckets uint,
		headAdditionalBuckets uint,
	) error
	Block(additional uint) error
	Link(additional uint) error
}
