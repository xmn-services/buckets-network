package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	stored_files "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents the storages factory
type Factory interface {
	Create() Factory
}

// Storages represents a storages instance
type Storages interface {
	entities.Mutable
	ToDownload() []files.File
	Stored() stored_files.Files
	Deleted() []files.File
	Download(file files.File) error
	Delete(file files.File) error
	Purge(hash hash.Hash) error
}
