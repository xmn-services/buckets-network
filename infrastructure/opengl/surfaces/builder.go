package surfaces

import (
	"errors"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/surfaces/surface"
)

type builder struct {
	surfaceBuilder surface.Builder
	renders        renders.Renders
	prog           programs.Program
}

func createBuilder(
	surfaceBuilder surface.Builder,
) Builder {
	out := builder{
		surfaceBuilder: surfaceBuilder,
		renders:        nil,
		prog:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.surfaceBuilder)
}

// WithRenders add renders to the builder
func (app *builder) WithRenders(renders renders.Renders) Builder {
	app.renders = renders
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(prog programs.Program) Builder {
	app.prog = prog
	return app
}

// Now builds a new Surfaces instance
func (app *builder) Now() (Surfaces, error) {
	if app.renders == nil {
		return nil, errors.New("the renders is mandatory in order to build a Surfaces instance")
	}

	out := []surface.Surface{}
	all := app.renders.All()
	for _, oneRender := range all {
		surfaceBuilder := app.surfaceBuilder.Create().WithRender(oneRender)
		if app.prog != nil {
			surfaceBuilder.WithProgram(app.prog)
		}

		surface, err := surfaceBuilder.Now()
		if err != nil {
			return nil, err
		}

		out = append(out, surface)
	}

	return createSurfaces(out), nil
}
