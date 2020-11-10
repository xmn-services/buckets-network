package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents a buckets factory
type Factory interface {
	Create() Buckets
}

// Buckets represents buckets
type Buckets interface {
	entities.Mutable
	All() []bucket.Bucket
	Add(bucket bucket.Bucket) error
	Delete(hash hash.Hash) error
	Fetch(absoluteFilePath string) (bucket.Bucket, error)
}
