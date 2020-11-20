package content

import (
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithContent(content []byte) Builder
	Now() (Content, error)
}

// Content represents a file content
type Content interface {
	Hash() hash.Hash
	Content() []byte
}
