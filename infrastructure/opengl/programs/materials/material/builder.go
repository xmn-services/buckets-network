package material

import (
	"errors"

	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	mat    *hash.Hash
	layers layers.Layers
}

func createBuilder() Builder {
	out := builder{
		mat:    nil,
		layers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMaterial adds a material to the builder
func (app *builder) WithMaterial(mat hash.Hash) Builder {
	app.mat = &mat
	return app
}

// WithCompiledLayers add compiled layers to the builder
func (app *builder) WithCompiledLayers(layers layers.Layers) Builder {
	app.layers = layers
	return app
}

// Now builds a new Material instance
func (app *builder) Now() (Material, error) {
	if app.mat == nil {
		return nil, errors.New("the material hash is mandatory in order to build a Material instance")
	}

	if app.layers == nil {
		return nil, errors.New("the compiled layers are mandatory in order to build a Material instance")
	}

	return createMaterial(*app.mat, app.layers), nil
}
