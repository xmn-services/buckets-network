package cameras

import (
	"time"

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

// Builder represents a camera builder
type Builder interface {
	Create() Builder
	WithViewport(viewport rectangles.Rectangle) Builder
	WithFieldOfView(fov float64) Builder
	WithIndex(index uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Camera, error)
}

// Camera represents a camera
type Camera interface {
	entities.Immutable
	Viewport() rectangles.Rectangle
	FieldOfView() float64
	Index() uint
}
