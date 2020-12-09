package textures

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/rectangles"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/rows"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents a texture builder
type Builder interface {
	Create() Builder
	WithViewport(viewport rectangles.Rectangle) Builder
	WithRows(rows rows.Rows) Builder
	Now() (Texture, error)
}

// Texture represents a texture
type Texture interface {
	entities.Immutable
	Viewport() rectangles.Rectangle
	Rows() rows.Rows
}
