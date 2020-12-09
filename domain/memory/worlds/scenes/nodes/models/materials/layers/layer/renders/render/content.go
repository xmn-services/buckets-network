package render

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/colors"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type content struct {
	tex    textures.Texture
	camera cameras.Camera
	color  colors.Color
}

func createContentWithTexture(
	tex textures.Texture,
) Content {
	return createContentInternally(tex, nil, nil)
}

func createContentWithCamera(
	camera cameras.Camera,
) Content {
	return createContentInternally(nil, camera, nil)
}

func createContentWithColor(
	color colors.Color,
) Content {
	return createContentInternally(nil, nil, color)
}

func createContentInternally(
	tex textures.Texture,
	camera cameras.Camera,
	color colors.Color,
) Content {
	out := content{
		tex:    tex,
		camera: camera,
		color:  color,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	if obj.IsTexture() {
		return obj.Texture().Hash()
	}

	return obj.Camera().Hash()
}

// IsTexture returns true if there is a texture, false otherwise
func (obj *content) IsTexture() bool {
	return obj.tex != nil
}

// Texture returns the texture, if any
func (obj *content) Texture() textures.Texture {
	return obj.tex
}

// IsCamera returns true if there is a camera, false otherwise
func (obj *content) IsCamera() bool {
	return obj.camera != nil
}

// Camera returns the camera, if any
func (obj *content) Camera() cameras.Camera {
	return obj.camera
}

// IsColor returns true if there is a color, false otherwise
func (obj *content) IsColor() bool {
	return obj.color != nil
}

// Color returns the color, if any
func (obj *content) Color() colors.Color {
	return obj.color
}
