package content

import (
	"github.com/xmn-services/buckets-network/libs/hash"
)

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
