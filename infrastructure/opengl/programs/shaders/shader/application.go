package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	domain "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders/shader"
)

type application struct {
	builder Builder
}

func createApplication(
	builder Builder,
) Application {
	out := application{
		builder: builder,
	}

	return &out
}

// Compile compiles a shader
func (app *application) Compile(shader domain.Shader) (Shader, error) {
	identifier, err := app.compileAny(shader)
	if err != nil {
		return nil, err
	}

	hsh := shader.Hash()
	return app.builder.Create().
		WithIdentifier(identifier).
		WithShader(hsh).
		Now()
}

func (app *application) compileAny(shader domain.Shader) (uint32, error) {
	code := shader.Code()
	if shader.Type().IsVertex() {
		return app.compile(code, gl.VERTEX_SHADER)
	}

	return app.compile(code, gl.FRAGMENT_SHADER)
}

func (app *application) compile(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
