package models

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	geo              geometries.Geometry
	mat              materials.Material
	variable         string
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		geo:              nil,
		mat:              nil,
		variable:         "",
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithGeometry adds a geometry to the builder
func (app *builder) WithGeometry(geo geometries.Geometry) Builder {
	app.geo = geo
	return app
}

// WithMaterial adds a material to the builder
func (app *builder) WithMaterial(mat materials.Material) Builder {
	app.mat = mat
	return app
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Model instance
func (app *builder) Now() (Model, error) {
	if app.geo == nil {
		return nil, errors.New("the geometry is mandatory in order to build a Model instance")
	}

	if app.mat == nil {
		return nil, errors.New("the material is mandatory in order to build a Model instance")
	}

	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Model instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.geo.Hash().Bytes(),
		app.mat.Hash().Bytes(),
		[]byte(app.variable),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createModel(immutable, app.geo, app.mat, app.variable), nil
}
