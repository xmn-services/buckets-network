package row

import "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels"

// Builder represents a row builder
type Builder interface {
	Create() Builder
	WithPixels(pixels pixels.Pixels) Builder
	Now() (Row, error)
}

// Row represents a row of pixels
type Row interface {
	Pixels() pixels.Pixels
	Length() int
}
