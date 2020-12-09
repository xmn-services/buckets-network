package materials

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type material struct {
	immutable entities.Immutable
	opacity   float64
	viewport  rectangles.Rectangle
	layers    layers.Layers
}

func createMaterial(
	immutable entities.Immutable,
	opacity float64,
	viewport rectangles.Rectangle,
	layers layers.Layers,
) Material {
	out := material{
		immutable: immutable,
		opacity:   opacity,
		viewport:  viewport,
		layers:    layers,
	}

	return &out
}

// Hash returns the hash
func (obj *material) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Opacity returns the opacity
func (obj *material) Opacity() float64 {
	return obj.opacity
}

// Viewport returns the viewport
func (obj *material) Viewport() rectangles.Rectangle {
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
