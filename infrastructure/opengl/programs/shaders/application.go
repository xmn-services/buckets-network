package shaders

import (
	domain_shaders "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders/shader"
)

type application struct {
	builder           Builder
	shaderApplication shader.Application
}

func createApplication(
	builder Builder,
	shaderApplication shader.Application,
) Application {
	out := application{
		builder:           builder,
		shaderApplication: shaderApplication,
	}

	return &out
}

// Compile compile shaders
func (app *application) Compile(shaders domain_shaders.Shaders) (Shaders, error) {
	compiledShadersList := []shader.Shader{}
	all := shaders.All()
	for _, oneShader := range all {
		compiledShader, err := app.shaderApplication.Compile(oneShader)
		if err != nil {
			return nil, err
		}

		compiledShadersList = append(compiledShadersList, compiledShader)
	}

	return app.builder.Create().WithCompiledShaders(compiledShadersList).Now()
}
