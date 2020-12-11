package materials

import (
	"errors"

	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials/material"
)

type builder struct {
	materials []material.Material
}

func createBuilder() Builder {
	out := builder{
		materials: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMaterials add materials to the builder
func (app *builder) WithMaterials(materials []material.Material) Builder {
	app.materials = materials
	return app
}

// Now builds a new Materials instance
func (app *builder) Now() (Materials, error) {
	if app.materials != nil && len(app.materials) <= 0 {
		app.materials = nil
	}

	if app.materials == nil {
		return nil, errors.New("the compiled materials are mandatory in order to build a Materials instance")
	}

	return createMaterials(app.materials), nil
}
