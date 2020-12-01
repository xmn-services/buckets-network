package accesses

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/accesses/access"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents an accesses factory
type Factory interface {
	Create() Accesses
}

// Accesses represents accesses
type Accesses interface {
	All() []access.Access
	Add(access access.Access) error
	Fetch(bucket hash.Hash) (access.Access, error)
	Delete(bucket hash.Hash) error
}
