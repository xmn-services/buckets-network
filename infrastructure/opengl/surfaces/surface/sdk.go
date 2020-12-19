package surface

import (
	image_color "image/color"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders/render"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/textures"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	textureBuilder := textures.NewBuilder()
	cameraBuilder := cameras.NewBuilder()
	return createBuilder(
		textureBuilder,
		cameraBuilder,
	)
}

// Builder represents a surface builder
type Builder interface {
	Create() Builder
	WithProgram(prog programs.Program) Builder
	WithRender(render render.Render) Builder
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
