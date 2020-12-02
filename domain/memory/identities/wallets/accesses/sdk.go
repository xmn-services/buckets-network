package accesses

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses/access"
	"github.com/xmn-services/buckets-network/libs/hash"
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
	return createBuilder()
}

// Adapter returns the accesses adapter
type Adapter interface {
	ToJSON(accesses Accesses) *JSONAccesses
	ToAccesses(ins *JSONAccesses) (Accesses, error)
}

// Factory represents an accesses factory
type Factory interface {
	Create() (Accesses, error)
}

// Builder represents an accesses builder
type Builder interface {
	Create() Builder
	WithList(lst []access.Access) Builder
	Now() (Accesses, error)
}

// Accesses represents accesses
type Accesses interface {
	All() []access.Access
	Add(access access.Access) error
	Fetch(bucket hash.Hash) (access.Access, error)
	Delete(bucket hash.Hash) error
}
