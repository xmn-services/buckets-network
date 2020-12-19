package applications

import (
	image_color "image/color"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/layers"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/nodes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/renders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/surfaces"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/textures"
)

type application struct {
	nodes nodes.Nodes
}

func createApplication(nodes nodes.Nodes) renders.Application {
	out := application{
		nodes: nodes,
	}

	return &out
}

// Render renders material on the texture
func (app *application) Render(mat materials.Material) (textures.Texture, error) {
	layers := mat.Layers()
	return app.layers(layers)
}

func (app *application) layers(layers layers.Layers) (textures.Texture, error) {
	all := layers.All()
	return app.layer(all[0])
}

func (app *application) layer(layer layer.Layer) (textures.Texture, error) {
	//alpha := layer.Alpha()
	viewport := layer.Viewport()
	surface := layer.Surface()

	tex, err := app.surface(surface, viewport)
	if err != nil {
		return nil, err
	}

	return tex, nil
}

func (app *application) surface(surface surfaces.Surface, viewport ints.Rectangle) (textures.Texture, error) {
	if surface.IsCamera() {
		cam := surface.Camera()
		return app.cameraToTexture(cam, viewport)
	}

	if surface.IsTexture() {
		return surface.Texture(), nil
	}

	color := surface.Color()
	return app.colorToTexture(color, viewport)
}

func (app *application) colorToTexture(color image_color.Color, viewport ints.Rectangle) (textures.Texture, error) {
	return nil, nil
}

func (app *application) cameraToTexture(cam cameras.Camera, viewport ints.Rectangle) (textures.Texture, error) {
	return nil, nil
}
