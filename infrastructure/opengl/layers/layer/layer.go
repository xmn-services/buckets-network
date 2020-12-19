package layer

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/surfaces"
)

type layer struct {
	alpha    uint8
	viewport ints.Rectangle
	surface  surfaces.Surface
}

func createLayer(
	alpha uint8,
	viewport ints.Rectangle,
	surface surfaces.Surface,
) Layer {
	out := layer{
		alpha:    alpha,
		viewport: viewport,
		surface:  surface,
	}

	return &out
}

// Alpha returns the alpha
func (obj *layer) Alpha() uint8 {
	return obj.alpha
}

// Viewport returns the viewport
func (obj *layer) Viewport() ints.Rectangle {
	return obj.viewport
}

// Surface returns the surface
func (obj *layer) Surface() surfaces.Surface {
	return obj.surface
}
