package files

import (
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents a files factory
type Factory interface {
	Create() Files
}

// Builder represents the files builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithoutHash() Builder
	WithFiles(hashes []hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Files, error)
}

// Files represents files
type Files interface {
	entities.Mutable
	All() []hash.Hash
	Exists(hash hash.Hash) bool
	Add(hash hash.Hash) error
	Delete(hash hash.Hash) error
}
