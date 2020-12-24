package materials

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/alphas"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/shaders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/viewports"
)

type material struct {
	id       *uuid.UUID
	alpha    alphas.Alpha
	shader   shaders.Shader
	viewport viewports.Viewport
	layers   []layers.Layer
}

func createMaterial(
	id *uuid.UUID,
	alpha alphas.Alpha,
	shader shaders.Shader,
	viewport viewports.Viewport,
	layers []layers.Layer,
) Material {
	out := material{
		id:       id,
		alpha:    alpha,
		shader:   shader,
		viewport: viewport,
		layers:   layers,
	}

	return &out
}

// ID returns the id
func (obj *material) ID() *uuid.UUID {
	return obj.id
}

// Alpha returns the alpha
func (obj *material) Alpha() alphas.Alpha {
	return obj.alpha
}

// Shader returns the shader
func (obj *material) Shader() shaders.Shader {
	return obj.shader
}

// Viewport returns the viewport
func (obj *material) Viewport() viewports.Viewport {
	return obj.viewport
}

// Layers returns the layers
func (obj *material) Layers() []layers.Layer {
	return obj.layers
}
