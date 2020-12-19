package programs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	domain_shaders "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/shaders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/shaders"
)

type builder struct {
	shadersBuilder shaders.Builder
	shaders        domain_shaders.Shaders
}

func createBuilder(
	shadersBuilder shaders.Builder,
) Builder {
	out := builder{
		shadersBuilder: shadersBuilder,
		shaders:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.shadersBuilder)
}

// WithShaders add shaders to the builder
func (app *builder) WithShaders(shaders domain_shaders.Shaders) Builder {
	app.shaders = shaders
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.shaders == nil {
		return nil, errors.New("the shaders are mandatory in order to build a Program instance")
	}

	// compile thwe shaders:
	shaders, err := app.shadersBuilder.Create().WithShaders(app.shaders).Now()
	if err != nil {
		return nil, err
	}

	// create program:
	program := gl.CreateProgram()

	// attach all compiled shaders:
	all := shaders.All()
	for _, oneShader := range all {
		gl.AttachShader(program, oneShader.Identifier())
	}

	// link the program:
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		str := fmt.Sprintf("failed to link program: %s", log)
		return nil, errors.New(str)
	}

	// delete compiled shaders:
	for _, oneShader := range all {
		gl.DeleteShader(oneShader.Identifier())
	}

	return createProgram(shaders, program), nil
}
