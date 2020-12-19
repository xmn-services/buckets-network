package materials

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
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
	WithAlpha(alpha uint8) Builder
	WithViewport(viewport ints.Rectangle) Builder
	WithLayers(layers layers.Layers) Builder
	WithShaders(shaders shaders.Shaders) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Material, error)
}

// Material represents a material
type Material interface {
	entities.Immutable
	Alpha() uint8
	Viewport() ints.Rectangle
	Layers() layers.Layers
	Shaders() shaders.Shaders
}
