package layer

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	domain_layer "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/surfaces"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	surfacesBuilder := surfaces.NewBuilder()
	return createBuilder(surfacesBuilder)
}

// Builder represents a layer builder
type Builder interface {
	Create() Builder
	WithLayer(layer domain_layer.Layer) Builder
	WithProgram(prog programs.Program) Builder
	Now() (Layer, error)
}

// Layer represents a layer of surfaces
type Layer interface {
	Alpha() uint8
	Viewport() ints.Rectangle
	Surfaces() surfaces.Surfaces
}
