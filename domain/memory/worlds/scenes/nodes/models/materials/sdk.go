package materials

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/alphas"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/shaders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/viewports"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the material builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithAlpha(alpha alphas.Alpha) Builder
	WithShader(shader shaders.Shader) Builder
	WithViewport(viewport viewports.Viewport) Builder
	WithLayers(layers []layers.Layer) Builder
	Now() (Material, error)
}

// Material represents a material
type Material interface {
	ID() *uuid.UUID
	Alpha() alphas.Alpha
	Shader() shaders.Shader
	Viewport() viewports.Viewport
	Layers() []layers.Layer
}
