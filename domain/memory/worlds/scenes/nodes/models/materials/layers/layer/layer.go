package layer

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type layer struct {
	immutable entities.Immutable
	opacity   float64
	viewport  rectangles.Rectangle
	renders   renders.Renders
	shaders   shaders.Shaders
}

func createLayer(
	immutable entities.Immutable,
	opacity float64,
	viewport rectangles.Rectangle,
	renders renders.Renders,
	shaders shaders.Shaders,
) Layer {
	out := layer{
		immutable: immutable,
		opacity:   opacity,
		viewport:  viewport,
		renders:   renders,
		shaders:   shaders,
	}

	return &out
}

// Hash returns the hash
func (obj *layer) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Opacity returns the opacity
func (obj *layer) Opacity() float64 {
	return obj.opacity
}

// Viewport returns the viewport
func (obj *layer) Viewport() rectangles.Rectangle {
	return obj.viewport
}

// Renders returns the renders
func (obj *layer) Renders() renders.Renders {
	return obj.renders
}

// Shaders returns the shaders
func (obj *layer) Shaders() shaders.Shaders {
	return obj.shaders
}

// CreatedOn returns the creation time
func (obj *layer) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
