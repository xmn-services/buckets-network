package layers

import (
	"errors"

	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers/layer"
)

type builder struct {
	list []layer.Layer
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithCompiledLayers add compiled layers to the builder
func (app *builder) WithCompiledLayers(layers []layer.Layer) Builder {
	app.list = layers
	return app
}

// Now builds a new Layers instance
func (app *builder) Now() (Layers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("the compiled layers are mandatory in order to build a Layers instance")
	}

	return createLayers(app.list), nil
}
