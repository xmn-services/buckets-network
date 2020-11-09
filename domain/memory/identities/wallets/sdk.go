package wallets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/follows"
)

// Factory represents a wallet factory
type Factory interface {
	Create() Wallet
}

// Wallet represents a wallet
type Wallet interface {
	Buckets() buckets.Buckets
	Follows() follows.Follow
}
