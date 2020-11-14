package files

import (
	stored_file "github.com/xmn-services/buckets-network/domain/memory/file"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents a files factory
type Factory interface {
	Create() Files
}

// Files represents files
type Files interface {
	entities.Mutable
	All() []stored_file.File
	Exists(hash hash.Hash) bool
	Add(file stored_file.File) error
	Fetch(hash hash.Hash) (stored_file.File, error)
	Delete(hash hash.Hash) error
}
