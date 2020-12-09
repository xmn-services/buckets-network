package cameras

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents a camera builder
type Builder interface {
	Create() Builder
	WithViewport(viewport rectangles.Rectangle) Builder
	WithFieldOfView(fov float64) Builder
	IsActive() Builder
	Now() (Camera, error)
}

// Camera represents a camera
type Camera interface {
	entities.Immutable
	Viewport() rectangles.Rectangle
	FieldOfView() float64
	IsActive() bool
}
