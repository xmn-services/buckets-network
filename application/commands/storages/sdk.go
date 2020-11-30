package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewApplication creates a new application instance
func NewApplication(
	bucketRepository buckets.Repository,
	contentRepository contents.Repository,
) Application {
	hashAdapter := hash.NewAdapter()
	return createApplication(hashAdapter, bucketRepository, contentRepository)
}

// Application represents a storage application
type Application interface {
	Exists(bucketHashStr string, chunkHashStr string) bool
	Retrieve(bucketHashStr string, chunkHashStr string) ([]byte, error)
}
