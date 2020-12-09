package geometries

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	shaders          shaders.Shaders
	vertices         vertices.Vertices
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		shaders:          nil,
		vertices:         nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithShaders add shaders to the builder
func (app *builder) WithShaders(shaders shaders.Shaders) Builder {
	app.shaders = shaders
	return app
}

// WithVertices add vertices to the builder
func (app *builder) WithVertices(vertices vertices.Vertices) Builder {
	app.vertices = vertices
	return app
}

// CreatedOn add creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Geometry instance
func (app *builder) Now() (Geometry, error) {
	if app.shaders == nil {
		return nil, errors.New("the shaders are mandatory in order to build a Geometry instance")
	}

	if app.vertices == nil {
		return nil, errors.New("the vertices are mandatory in order to build a Geometry instance")
	}

	if app.shaders.IsVertex() {
		return nil, errors.New("the geometry shaders were expected to be vertex shaders")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.shaders.Hash().Bytes(),
		app.vertices.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createGeometry(immutable, app.shaders, app.vertices), nil
}
