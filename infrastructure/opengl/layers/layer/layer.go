package layer

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/surfaces"
)

type layer struct {
	alpha    uint8
	viewport ints.Rectangle
	surfaces surfaces.Surfaces
}

func createLayer(
	alpha uint8,
	viewport ints.Rectangle,
	surfaces surfaces.Surfaces,
) Layer {
	out := layer{
		alpha:    alpha,
		viewport: viewport,
		surfaces: surfaces,
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

// Surfaces returns the surfaces
func (obj *layer) Surfaces() surfaces.Surfaces {
	return obj.surfaces
}
