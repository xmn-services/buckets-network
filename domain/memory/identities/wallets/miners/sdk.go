package miners

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/permanents"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new miner instance
func NewBuilder() Builder {
	blocksFactory := blocks.NewFactory()
	bucketsFactory := buckets.NewFactory()
	pBucketsFactory := permanents.NewFactory()
	return createBuilder(blocksFactory, bucketsFactory, pBucketsFactory)
}

// Adapter represents the miner adapter
type Adapter interface {
	ToJSON(ins Miner) *JSONMiner
	ToMiner(js *JSONMiner) (Miner, error)
}

// Factory represents a miner factory
type Factory interface {
	Create() (Miner, error)
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithToTransact(toTransact buckets.Buckets) Builder
	WithQueue(queue buckets.Buckets) Builder
	WithBroadcasted(broadcasted permanents.Buckets) Builder
	WithToLink(toLink blocks.Blocks) Builder
	Now() (Miner, error)
}

// Miner represents the miner
type Miner interface {
	ToTransact() buckets.Buckets
	Queue() buckets.Buckets
	Broadcasted() permanents.Buckets
	ToLink() blocks.Blocks
}
