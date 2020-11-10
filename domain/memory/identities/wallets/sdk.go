package wallets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/follows"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/permanents"
)

// Factory represents a wallet factory
type Factory interface {
	Create() Wallet
}

// Wallet represents a wallet
type Wallet interface {
	All() permanents.Permanent
	New() buckets.Buckets
	Queue() buckets.Buckets
	Follows() follows.Follow
	Broadcasted() permanents.Permanent
	Blocks() blocks.Blocks
}
