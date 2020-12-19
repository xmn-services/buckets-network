package surfaces

import (
	image_color "image/color"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/textures"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	textureBuilder := textures.NewBuilder()
	return createBuilder(
		textureBuilder,
	)
}

// Builder represents a surface builder
type Builder interface {
	Create() Builder
	WithProgram(prog programs.Program) Builder
	WithRender(render renders.Render) Builder
	Now() (Surface, error)
}

// Surface represents a surface
type Surface interface {
	IsCamera() bool
	Camera() cameras.Camera
	IsTexture() bool
	Texture() textures.Texture
	IsColor() bool
	Color() image_color.Color
}
