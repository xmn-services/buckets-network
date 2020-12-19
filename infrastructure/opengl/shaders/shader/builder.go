package shader

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	domain_shader "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/shaders/shader"
)

type builder struct {
	shader domain_shader.Shader
}

func createBuilder() Builder {
	out := builder{
		shader: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithShader adds a shader to the builder
func (app *builder) WithShader(shader domain_shader.Shader) Builder {
	app.shader = shader
	return app
}

// Now builds a new Shader instance
func (app *builder) Now() (Shader, error) {

	if app.shader == nil {
		return nil, errors.New("the shader is mandatory in order to build a Shader instance")
	}

	identifier, err := app.compileAny(app.shader)
	if err != nil {
		return nil, err
	}

	return createShader(app.shader, identifier), nil
}

func (app *builder) compileAny(shader domain_shader.Shader) (uint32, error) {
	code := shader.Code()
	if shader.Type().IsVertex() {
		return app.compile(code, gl.VERTEX_SHADER)
	}

	return app.compile(code, gl.FRAGMENT_SHADER)
}

func (app *builder) compile(source string, shaderType uint32) (uint32, error) {
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
