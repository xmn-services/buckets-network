package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/buckets/bucket"
)

// Factory represents a buckets factory
type Factory interface {
	Create() Buckets
}

// Buckets represents buckets
type Buckets interface {
	All() []bucket.Bucket
	Add(bucket bucket.Bucket) error
	Delete(bucket bucket.Bucket) error
	Fetch(absoluteFilePath string) (bucket.Bucket, error)
}
