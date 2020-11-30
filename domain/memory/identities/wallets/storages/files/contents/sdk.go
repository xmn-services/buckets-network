package contents

import (
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the content adapter
type Adapter interface {
	ToJSON(content Content) *JSONContent
	ToContent(js *JSONContent) (Content, error)
}

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithBucket(bucket hash.Hash) Builder
	WithFile(file hash.Hash) Builder
	WithChunk(chunk hash.Hash) Builder
	Now() (Content, error)
}

// Content represents contents
type Content interface {
	Bucket() hash.Hash
	File() hash.Hash
	Chunk() hash.Hash
}
