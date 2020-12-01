package access

import (
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

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
