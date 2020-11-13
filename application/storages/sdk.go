package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	stored_file "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files"
)

// Application represents a storage application
type Application interface {
	Exists(file files.File) bool
	Retrieve(file files.File) (stored_file.File, error)
	Delete(file files.File) error
}
