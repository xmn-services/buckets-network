package storages

import (
	stored_file "github.com/xmn-services/buckets-network/domain/memory/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewApplication creates a new application instance
func NewApplication(
	storedFileRepository stored_file.Repository,
) Application {
	hashAdapter := hash.NewAdapter()
	return createApplication(hashAdapter, storedFileRepository)
}

// Application represents a storage application
type Application interface {
	IsStored(fileHashStr string) bool
	Retrieve(fileHashStr string) (stored_file.File, error)
}
