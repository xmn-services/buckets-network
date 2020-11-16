package wallets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages"
)

// Factory represents a wallet factory
type Factory interface {
	Create() Wallet
}

// Wallet represents a wallet
type Wallet interface {
	Miner() miners.Miner
	Storage() storages.Storage
}
