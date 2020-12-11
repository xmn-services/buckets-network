package layer

import (
	"errors"

	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	layer   *hash.Hash
	shaders shaders.Shaders
}

func createBuilder() Builder {
	out := builder{
		layer:   nil,
		shaders: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLayer adds a layer hash to the builder
func (app *builder) WithLayer(layer hash.Hash) Builder {
	app.layer = &layer
	return app
}

// WithCompiledShaders add compiled shaders to the builder
func (app *builder) WithCompiledShaders(shaders shaders.Shaders) Builder {
	app.shaders = shaders
	return app
}

// Now builds a new Layer instance
func (app *builder) Now() (Layer, error) {
	if app.layer == nil {
		return nil, errors.New("the layer hash is mandatory in order to build a Layer instance")
	}

	if app.shaders == nil {
		return nil, errors.New("the compiled shaders are mandatory in order to build a Layer instance")
	}

	return createLayer(*app.layer, app.shaders), nil
}
