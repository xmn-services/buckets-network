package layers

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders/shader"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a layers builder
type Builder interface {
	Create() Builder
	WithCompiledLayers(layers []layer.Layer) Builder
	Now() (Layers, error)
}

// Layers represents layers
type Layers interface {
	All() []layer.Layer
	CompiledShaders() []shader.Shader
}
