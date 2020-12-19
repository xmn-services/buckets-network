package surface

import (
	image_color "image/color"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/textures"
)

type surface struct {
	cam cameras.Camera
	tex textures.Texture
	col image_color.Color
}

func createSurfaceWithCamera(
	cam cameras.Camera,
) Surface {
	return createSurfaceInternally(cam, nil, nil)
}

func createSurfaceWithTexture(
	tex textures.Texture,
) Surface {
	return createSurfaceInternally(nil, tex, nil)
}

func createSurfaceWithColor(
	col image_color.Color,
) Surface {
	return createSurfaceInternally(nil, nil, col)
}

func createSurfaceInternally(
	cam cameras.Camera,
	tex textures.Texture,
	col image_color.Color,
) Surface {
	out := surface{
		cam: cam,
		tex: tex,
		col: col,
	}

	return &out
}

// IsCamera returns true if there is a camera, false otherwise
func (obj *surface) IsCamera() bool {
	return obj.cam != nil
}

// Camera returns the camera, if any
func (obj *surface) Camera() cameras.Camera {
	return obj.cam
}

// IsTexture returns true if there is a texture, false otherwise
func (obj *surface) IsTexture() bool {
	return obj.tex != nil
}

// Texture returns the texture, if any
func (obj *surface) Texture() textures.Texture {
	return obj.tex
}

// IsColor returns true if there is a color, false otherwise
func (obj *surface) IsColor() bool {
	return obj.col != nil
}

// Color returns the color, if any
func (obj *surface) Color() image_color.Color {
	return obj.col
}
