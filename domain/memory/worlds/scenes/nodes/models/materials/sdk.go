package materials

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Builder represents the material builder
type Builder interface {
	Create() Builder
	WithOpacity(opacity float64) Builder
	WithViewport(viewport rectangles.Rectangle) Builder
	WithLayers(layers layers.Layers) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Material, error)
}

// Material represents a material
type Material interface {
	entities.Immutable
	Opacity() float64
	Viewport() rectangles.Rectangle
	Layers() layers.Layers
}
