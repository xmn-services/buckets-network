package layer

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
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
	WithAlpha(alpha uint8) Builder
	WithViewport(viewport ints.Rectangle) Builder
	WithRender(render renders.Render) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Layer, error)
}

// Layer represents layer of textures
type Layer interface {
	entities.Immutable
	Alpha() uint8
	Viewport() ints.Rectangle
	Render() renders.Render
}
