package layer

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a layer builder
type Builder interface {
	Create() Builder
	WithLayer(layer hash.Hash) Builder
	WithCompiledShaders(shaders shaders.Shaders) Builder
	Now() (Layer, error)
}

// Layer represents a layer
type Layer interface {
	Layer() hash.Hash
	Compiled() shaders.Shaders
}
