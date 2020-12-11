package shaders

import (
	"errors"

	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders/shader"
)

type builder struct {
	list []shader.Shader
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

// WithCompiledShaders add shaders to the builder
func (app *builder) WithCompiledShaders(shaders []shader.Shader) Builder {
	app.list = shaders
	return app
}

// Now builds a new Shaders instance
func (app *builder) Now() (Shaders, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("the compiled shaders are mandatory in order to build a Shaders instance")
	}

	return createShaders(app.list), nil
}
