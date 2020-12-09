package layer

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
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

// Builder represents a layer builder
type Builder interface {
	Create() Builder
	WithOpacity(opacity float64) Builder
	WithViewport(viewport rectangles.Rectangle) Builder
	WithRenders(renders renders.Renders) Builder
	WithShaders(shaders shaders.Shaders) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Layer, error)
}

// Layer represents layer of textures
type Layer interface {
	entities.Immutable
	Opacity() float64
	Viewport() rectangles.Rectangle
	Renders() renders.Renders
	Shaders() shaders.Shaders
}
