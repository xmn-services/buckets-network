package materials

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/layers"
)

type material struct {
	alpha    uint8
	viewport ints.Rectangle
	layers   layers.Layers
}

func createMaterial(
	alpha uint8,
	viewport ints.Rectangle,
	layers layers.Layers,
) Material {
	out := material{
		alpha:    alpha,
		viewport: viewport,
		layers:   layers,
	}

	return &out
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

// Render renders a material
func (obj *material) Render() error {
	return nil
}
