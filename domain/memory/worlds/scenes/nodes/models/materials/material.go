package materials

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type material struct {
	immutable entities.Immutable
	alpha     uint8
	viewport  ints.Rectangle
	layers    layers.Layers
}

func createMaterial(
	immutable entities.Immutable,
	alpha uint8,
	viewport ints.Rectangle,
	layers layers.Layers,
) Material {
	out := material{
		immutable: immutable,
		alpha:     alpha,
		viewport:  viewport,
		layers:    layers,
	}

	return &out
}

// Hash returns the hash
func (obj *material) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Alpha returns the alpha
func (obj *material) Alpha() uint8 {
	return obj.alpha
}

// Viewport returns the viewport
func (obj *material) Viewport() ints.Rectangle {
	return obj.viewport
}

// Layers returns the layers
func (obj *material) Layers() layers.Layers {
	return obj.layers
}

// CreatedOn returns the creation time
func (obj *material) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
