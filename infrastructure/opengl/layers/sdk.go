package layers

import (
	domain_layers "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	layerBuilder := layer.NewBuilder()
	return createBuilder(layerBuilder)
}

// Builder represents a layers builder
type Builder interface {
	Create() Builder
	WithLayers(layers domain_layers.Layers) Builder
	WithProgram(prog programs.Program) Builder
	Now() (Layers, error)
}

// Layers represents layers
type Layers interface {
	All() []layer.Layer
}
