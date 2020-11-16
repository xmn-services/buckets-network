package wallets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages"
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

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	minerFactory := miners.NewFactory()
	storageFactory := storages.NewFactory()
	return createBuilder(minerFactory, storageFactory)
}

// Adapter represents the wallet adapter
type Adapter interface {
	ToJSON(ins Wallet) *JSONWallet
	ToWallet(js *JSONWallet) (Wallet, error)
}

// Factory represents a wallet factory
type Factory interface {
	Create() (Wallet, error)
}

// Builder represents a wallet builder
type Builder interface {
	Create() Builder
	WithMiner(miner miners.Miner) Builder
	WithStorage(storage storages.Storage) Builder
	Now() (Wallet, error)
}

// Wallet represents a wallet
type Wallet interface {
	Miner() miners.Miner
	Storage() storages.Storage
}
