package program

import (
	"errors"

	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials"
)

type builder struct {
	materials  materials.Materials
	identifier uint32
}

func createBuilder() Builder {
	out := builder{
		materials:  nil,
		identifier: uint32(0),
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithCompiledMaterials add compiled materials to the builder
func (app *builder) WithCompiledMaterials(materials materials.Materials) Builder {
	app.materials = materials
	return app
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier uint32) Builder {
	app.identifier = identifier
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.materials == nil {
		return nil, errors.New("the materials are mandatory in order to build a Program instance")
	}

	if app.identifier == 0 {
		return nil, errors.New("the identifier is mandatory in order to build a Program instance")
	}

	return createProgram(app.materials, app.identifier), nil
}
