package materials

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the material builder
type Builder interface {
	Create() Builder
	WithOpacity(opacity float64) Builder
	WithViewport(viewport rectangles.Rectangle) Builder
	WithLayers(layers layers.Layers) Builder
	Now() (Material, error)
}

// Material represents a material
type Material interface {
	entities.Immutable
	Opacity() float64
	Viewport() rectangles.Rectangle
	Layers() layers.Layers
}
