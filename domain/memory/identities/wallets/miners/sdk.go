package miners

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/permanents"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Factory represents a miner factory
type Factory interface {
	Create() Miner
}

// Miner represents the miner
type Miner interface {
	entities.Mutable
	ToTransact() buckets.Buckets
	Queue() buckets.Buckets
	Broadcasted() permanents.Buckets
	Blocks() blocks.Blocks
}
