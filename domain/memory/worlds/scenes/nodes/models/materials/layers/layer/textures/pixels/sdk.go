package pixels

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels/pixel"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the pixels builder
type Builder interface {
	Create() Builder
	WithPixels(pixels []pixel.Pixel) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Pixels, error)
}

// Pixels represents pixels
type Pixels interface {
	entities.Immutable
	All() []pixel.Pixel
	Amount() uint
}
