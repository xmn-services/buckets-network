package layer

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type layer struct {
	immutable entities.Immutable
	alpha     uint8
	viewport  ints.Rectangle
	render    renders.Render
}

func createLayer(
	immutable entities.Immutable,
	alpha uint8,
	viewport ints.Rectangle,
	render renders.Render,
) Layer {
	out := layer{
		immutable: immutable,
		alpha:     alpha,
		viewport:  viewport,
		render:    render,
	}

	return &out
}

// Hash returns the hash
func (obj *layer) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Alpha returns the alpha
func (obj *layer) Alpha() uint8 {
	return obj.alpha
}

// Viewport returns the viewport
func (obj *layer) Viewport() ints.Rectangle {
	return obj.viewport
}

// Render returns the render
func (obj *layer) Render() renders.Render {
	return obj.render
}

// CreatedOn returns the creation time
func (obj *layer) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
