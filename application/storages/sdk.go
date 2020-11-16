package storages

import (
	stored_file "github.com/xmn-services/buckets-network/domain/memory/file"
)

// Application represents a storage application
type Application interface {
	IsStored(fileHashStr string) bool
	Retrieve(fileHashStr string) (stored_file.File, error)
}
