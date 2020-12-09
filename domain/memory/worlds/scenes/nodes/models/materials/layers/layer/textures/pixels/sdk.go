package pixels

import "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels/pixel"

// Builder represents the pixels builder
type Builder interface {
	Create() Builder
	WithPixels(pixels []pixel.Pixel) Builder
	Now() (Pixels, error)
}

// Pixels represents pixels
type Pixels interface {
	All() []pixel.Pixel
}
