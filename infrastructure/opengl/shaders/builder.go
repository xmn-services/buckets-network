package shaders

import (
	"errors"

	domain_shaders "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/shaders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/shaders/shader"
)

type builder struct {
	shaderBuilder shader.Builder
	domainShaders domain_shaders.Shaders
}

func createBuilder(
	shaderBuilder shader.Builder,
) Builder {
	out := builder{
		shaderBuilder: shaderBuilder,
		domainShaders: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.shaderBuilder)
}

// WithShaders add shaders to the list
func (app *builder) WithShaders(shaders domain_shaders.Shaders) Builder {
	app.domainShaders = shaders
	return app
}

// Now builds a new Shaders instance
func (app *builder) Now() (Shaders, error) {
	if app.domainShaders == nil {
		return nil, errors.New("the shaders are mandatory in order to build a Shaders instance")
	}

	list := []shader.Shader{}
	all := app.domainShaders.All()
	for _, oneDomainShader := range all {
		shader, err := app.shaderBuilder.Create().WithShader(oneDomainShader).Now()
		if err != nil {
			return nil, err
		}

		list = append(list, shader)
	}

	return createShaders(list), nil
}
