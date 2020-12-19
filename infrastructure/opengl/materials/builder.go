package materials

import (
	"errors"

	domain_materials "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/layers"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

type builder struct {
	layersBuilder layers.Builder
	material      domain_materials.Material
	prog          programs.Program
}

func createBuilder(
	layersBuilder layers.Builder,
) Builder {
	out := builder{
		layersBuilder: layersBuilder,
		material:      nil,
		prog:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.layersBuilder)
}

// WithMaterial adds a domain material to the builder
func (app *builder) WithMaterial(material domain_materials.Material) Builder {
	app.material = material
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(prog programs.Program) Builder {
	app.prog = prog
	return app
}

// Now builds a new Material instance
func (app *builder) Now() (Material, error) {
	if app.material == nil {
		return nil, errors.New("the material is mandatory in order to build a Material instance")
	}

	domainLayers := app.material.Layers()
	layersBuilder := app.layersBuilder.Create().WithLayers(domainLayers)
	if app.prog != nil {
		layersBuilder.WithProgram(app.prog)
	}

	layers, err := layersBuilder.Now()
	if err != nil {
		return nil, err
	}

	alpha := app.material.Alpha()
	viewport := app.material.Viewport()
	return createMaterial(alpha, viewport, layers), nil
}
