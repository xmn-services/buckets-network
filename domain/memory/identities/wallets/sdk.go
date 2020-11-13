package wallets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Factory represents a wallet factory
type Factory interface {
	Create() Wallet
}

// Wallet represents a wallet
type Wallet interface {
	entities.Mutable
	Miner() miners.Miner
	Storages() storages.Storages
}
