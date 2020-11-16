package files

import (
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the files adapter
type Adapter interface {
	ToJSON(ins Files) *JSONFiles
	ToFiles(js *JSONFiles) (Files, error)
}

// Factory represents a files factory
type Factory interface {
	Create() (Files, error)
}

// Builder represents the files builder
type Builder interface {
	Create() Builder
	WithFiles(hashes []hash.Hash) Builder
	Now() (Files, error)
}

// Files represents files
type Files interface {
	All() []hash.Hash
	Exists(hash hash.Hash) bool
	Add(hash hash.Hash) error
	Delete(hash hash.Hash) error
}
