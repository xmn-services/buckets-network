package layer

import (
	"errors"

	domain_layer "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/surfaces"
)

type builder struct {
	surfaceBuilder surfaces.Builder
	layer          domain_layer.Layer
	prog           programs.Program
}

func createBuilder(
	surfaceBuilder surfaces.Builder,
) Builder {
	out := builder{
		surfaceBuilder: surfaceBuilder,
		layer:          nil,
		prog:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.surfaceBuilder)
}

// WithLayer adds a layer to the builder
func (app *builder) WithLayer(layer domain_layer.Layer) Builder {
	app.layer = layer
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(prog programs.Program) Builder {
	app.prog = prog
	return app
}

// Now builds a new Layer instance
func (app *builder) Now() (Layer, error) {
	if app.layer == nil {
		return nil, errors.New("the layer is mandatory in order to build a Layer instance")
	}

	render := app.layer.Render()
	surfaceBuilder := app.surfaceBuilder.Create().WithRender(render)
	if app.prog != nil {
		surfaceBuilder.WithProgram(app.prog)
	}

	surface, err := surfaceBuilder.Now()
	if err != nil {
		return nil, err
	}

	alpha := app.layer.Alpha()
	viewport := app.layer.Viewport()
	return createLayer(alpha, viewport, surface), nil
}
