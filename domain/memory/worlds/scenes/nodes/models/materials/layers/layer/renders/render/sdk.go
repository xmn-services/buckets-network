package render

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/colors"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/rectangles"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents a render builder
type Builder interface {
	Create() Builder
	WithOpacity(opacity float64) Builder
	WithBackground(color colors.Color) Builder
	WithViewport(viewport rectangles.Rectangle) Builder
	WithTexture(tex textures.Texture) Builder
	WithCamera(cam cameras.Camera) Builder
	Now() (Render, error)
}

// Render represents layer render
type Render interface {
	entities.Immutable
	Opacity() float64
	Background() colors.Color
	Viewport() rectangles.Rectangle
	Content() Content
}

// Content represents a render content
type Content interface {
	IsTexture() bool
	Texture() textures.Texture
	IsCamera() bool
	Camera() cameras.Camera
}
