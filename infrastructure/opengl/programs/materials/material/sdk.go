package material

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a material builder
type Builder interface {
	Create() Builder
	WithMaterial(mat hash.Hash) Builder
	WithCompiledLayers(layers layers.Layers) Builder
	Now() (Material, error)
}

// Material represents a material
type Material interface {
	Material() hash.Hash
	Compiled() layers.Layers
}
