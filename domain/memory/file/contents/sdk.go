package contents

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/file/contents/content"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	contentBuilder := content.NewBuilder()
	return createBuilder(contentBuilder)
}

// Builder represents a contents builder
type Builder interface {
	Create() Builder
	WithContents(contents [][]byte) Builder
	Now() (Contents, error)
}

// Contents represents a content
type Contents interface {
	All() []content.Content
	Add(content content.Content) error
	NotStored(file files.File) uint
}
