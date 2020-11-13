package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	stored_file "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Storages represents a storages instance
type Storages interface {
	entities.Mutable
	ToDownload() []files.File
	Stored() []stored_file.File
	Deleted() []files.File
	Download(file files.File) error
	Store(file stored_file.File) error
	Delete(file files.File) error
	Purge(hash hash.Hash) error
}
