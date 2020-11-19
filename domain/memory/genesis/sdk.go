package genesis

import (
	"time"

	transfer_genesis "github.com/xmn-services/buckets-network/domain/transfers/genesis"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	trService transfer_genesis.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	trRepository transfer_genesis.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(builder, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_genesis.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the genesis adapter
type Adapter interface {
	ToTransfer(genesis Genesis) (transfer_genesis.Genesis, error)
	ToJSON(genesis Genesis) *JSONGenesis
	ToGenesis(ins *JSONGenesis) (Genesis, error)
}

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithMiningValue(miningValue uint8) Builder
	WithBlockDifficultyBase(blockDiffBase uint) Builder
	WithBlockDifficultyIncreasePerBucket(blockDiffIncreasePerBucket float64) Builder
	WithLinkDifficulty(link uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents the genesis
type Genesis interface {
	entities.Immutable
	MiningValue() uint8
	Difficulty() Difficulty
}

// Difficulty represents the genesis difficulty
type Difficulty interface {
	Block() Block
	Link() uint
}

// Block represents the block difficulty related data
type Block interface {
	Base() uint
	IncreasePerBucket() float64
}

// Repository repreents the genesis repository
type Repository interface {
	Retrieve() (Genesis, error)
}

// Service represents the genesis service
type Service interface {
	Save(genesis Genesis) error
}
