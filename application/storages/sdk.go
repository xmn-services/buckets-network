package storages

import (
	stored_file "github.com/xmn-services/buckets-network/domain/memory/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Application represents a storage application
type Application interface {
	IsStored(file hash.Hash) bool
	Retrieve(file hash.Hash) (stored_file.File, error)
}
