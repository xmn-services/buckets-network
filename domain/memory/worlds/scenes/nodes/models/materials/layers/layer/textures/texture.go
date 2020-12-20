package textures

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type texture struct {
	immutable entities.Immutable
	viewport  ints.Rectangle
	pixels    []pixels.Pixel
}

func createTexture(
	immutable entities.Immutable,
	viewport ints.Rectangle,
	pixels []pixels.Pixel,
) Texture {
	out := texture{
		immutable: immutable,
		viewport:  viewport,
		pixels:    pixels,
	}

	return &out
}

// Hash returns the hash
func (obj *texture) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Viewport returns the viewport
func (obj *texture) Viewport() ints.Rectangle {
	return obj.viewport
}

// Pixels returns the pixels
func (obj *texture) Pixels() []pixels.Pixel {
	return obj.pixels
}

// CreatedOn returns the creation time
func (obj *texture) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
