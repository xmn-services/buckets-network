package chains

import (
	"time"

	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	transfer_chains "github.com/xmn-services/buckets-network/domain/transfers/chains"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	genesisService genesis.Service,
	blockRepository mined_block.Repository,
	blockService mined_block.Service,
	linkRepository mined_link.Repository,
	linkService mined_link.Service,
	trService transfer_chains.Service,
) Service {
	hashAdapter := hash.NewAdapter()
	adapter := NewAdapter()
	return createService(hashAdapter, adapter, repository, genesisService, blockRepository, blockService, linkRepository, linkService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	genesisRepository genesis.Repository,
	blockRepository mined_block.Repository,
	linkRepository mined_link.Repository,
	trRepository transfer_chains.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(genesisRepository, blockRepository, linkRepository, trRepository, builder)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_chains.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// BlockDifficulty calculates the block difficulty
func BlockDifficulty(baseDifficulty uint, increasePerBucket float64, amountBuckets uint) uint {
	sum := float64(0)
	base := float64(baseDifficulty)
	for i := 0; i < int(amountBuckets); i++ {
		sum += increasePerBucket
	}

	return uint(sum + base)
}

// Adapter returns the chain adapter
type Adapter interface {
	JSONToChain(js []byte) (Chain, error)
	ToTransfer(chain Chain) (transfer_chains.Chain, error)
}

// Builder represents the chain builder
type Builder interface {
	Create() Builder
	WithGenesis(gen genesis.Genesis) Builder
	WithRoot(root mined_block.Block) Builder
	WithHead(head mined_link.Link) Builder
	WithTotal(total uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Chain, error)
}

// Chain represents a chain
type Chain interface {
	entities.Immutable
	Genesis() genesis.Genesis
	Root() mined_block.Block
	Head() mined_link.Link
	Total() uint
	Height() uint
}

// Repository represents the chain repository
type Repository interface {
	Retrieve() (Chain, error)
	RetrieveAtIndex(index uint) (Chain, error)
}

// Service represents the chain service
type Service interface {
	Insert(chain Chain) error
	Update(original Chain, updated Chain) error
}
