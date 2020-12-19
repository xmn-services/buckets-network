package materials

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	domain_materials "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/layers"
)

type material struct {
	original domain_materials.Material
	alpha    uint8
	viewport ints.Rectangle
	layers   layers.Layers
}

func createMaterial(
	original domain_materials.Material,
	alpha uint8,
	viewport ints.Rectangle,
	layers layers.Layers,
) Material {
	out := material{
		original: original,
		alpha:    alpha,
		viewport: viewport,
		layers:   layers,
	}

	return &out
}

// Original returns the original material
func (obj *material) Original() domain_materials.Material {
	return obj.original
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
