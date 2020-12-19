package geometries

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter           hash.Adapter
	immutableBuilder      entities.ImmutableBuilder
	verticesFactory       vertices.Factory
	vertices              vertices.Vertices
	vertexCoordinatesVar  string
	textureCoordinatesVar string
	createdOn             *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
	verticesFactory vertices.Factory,
) Builder {
	out := builder{
		hashAdapter:           hashAdapter,
		immutableBuilder:      immutableBuilder,
		verticesFactory:       verticesFactory,
		vertices:              nil,
		vertexCoordinatesVar:  "",
		textureCoordinatesVar: "",
		createdOn:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder, app.verticesFactory)
}

// WithVertices add vertices to the builder
func (app *builder) WithVertices(vertices vertices.Vertices) Builder {
	app.vertices = vertices
	return app
}

// WithVertexCoordinatesVariable adds a vertex coordinates variable to the builder
func (app *builder) WithVertexCoordinatesVariable(vertexCoordinatesVar string) Builder {
	app.vertexCoordinatesVar = vertexCoordinatesVar
	return app
}

// WithTextureCoordinatesVariable adds a texture coordinates variable to the builder
func (app *builder) WithTextureCoordinatesVariable(texCoordinatesVar string) Builder {
	app.textureCoordinatesVar = texCoordinatesVar
	return app
}

// CreatedOn add creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Geometry instance
func (app *builder) Now() (Geometry, error) {
	if app.vertices == nil {
		vertices, err := app.verticesFactory.Create()
		if err != nil {
			return nil, err
		}

		app.vertices = vertices
	}

	if app.vertexCoordinatesVar == "" {
		return nil, errors.New("the vertex coordinates variable is mandatory in order to build a Geometry instance")
	}

	if app.textureCoordinatesVar == "" {
		return nil, errors.New("the texture coordinates variable is mandatory in order to build a Geometry instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.vertices.Hash().Bytes(),
		[]byte(app.vertexCoordinatesVar),
		[]byte(app.textureCoordinatesVar),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	variables := createVariables(app.vertexCoordinatesVar, app.textureCoordinatesVar)
	return createGeometry(immutable, variables, app.vertices), nil
}
