package access

import (
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter returns the access adapter
type Adapter interface {
	ToJSON(access Access) *JSONAccess
	ToAccess(ins *JSONAccess) (Access, error)
}

// Builder represents an access builder
type Builder interface {
	Create() Builder
	WithBucket(bucket hash.Hash) Builder
	WithKey(key public.Key) Builder
	Now() (Access, error)
}

// Access represents an access
type Access interface {
	Bucket() hash.Hash
	Key() public.Key
}
