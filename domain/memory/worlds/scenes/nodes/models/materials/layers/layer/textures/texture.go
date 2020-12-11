package textures

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/rows"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type texture struct {
	immutable entities.Immutable
	dimension rectangles.Rectangle
	pixels    rows.Rows
}

func createTexture(
	immutable entities.Immutable,
	dimension rectangles.Rectangle,
	pixels rows.Rows,
) Texture {
	out := texture{
		immutable: immutable,
		dimension: dimension,
		pixels:    pixels,
	}

	return &out
}

// Hash returns the hash
func (obj *texture) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Dimension returns the dimension
func (obj *texture) Dimension() rectangles.Rectangle {
	return obj.dimension
}

// Pixels returns the pixels
func (obj *texture) Pixels() rows.Rows {
	return obj.pixels
}

// CreatedOn returns the creation time
func (obj *texture) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
