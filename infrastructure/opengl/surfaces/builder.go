package surfaces

import (
	"errors"
	image_color "image/color"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/textures"
)

type builder struct {
	textureBuilder textures.Builder
	render         renders.Render
	prog           programs.Program
}

func createBuilder(
	textureBuilder textures.Builder,
) Builder {
	out := builder{
		textureBuilder: textureBuilder,
		render:         nil,
		prog:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.textureBuilder)
}

// WithRender adds a render to the builder
func (app *builder) WithRender(render renders.Render) Builder {
	app.render = render
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(prog programs.Program) Builder {
	app.prog = prog
	return app
}

// Now builds a new Surface instance
func (app *builder) Now() (Surface, error) {
	if app.render == nil {
		return nil, errors.New("the render is mandatory in order to build a Surface instance")
	}

	content := app.render.Content()
	if content.IsTexture() {
		domainTex := content.Texture()
		tex, err := app.textureBuilder.Create().WithTexture(domainTex).Now()
		if err != nil {
			return nil, err
		}

		return createSurfaceWithTexture(tex), nil
	}

	if content.IsCamera() {
		cam := content.Camera()
		return createSurfaceWithCamera(cam), nil
	}

	col := content.Color()
	rgba := image_color.RGBA{
		R: col.Red(),
		G: col.Green(),
		B: col.Blue(),
		A: 0xff,
	}

	return createSurfaceWithColor(rgba), nil
}
