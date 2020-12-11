package shader

import (
	"errors"

	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	sh *hash.Hash
	id uint32
}

func createBuilder() Builder {
	out := builder{
		sh: nil,
		id: uint32(0),
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithShader adds a shader to the builder
func (app *builder) WithShader(shader hash.Hash) Builder {
	app.sh = &shader
	return app
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier uint32) Builder {
	app.id = identifier
	return app
}

// Now builds a new Shader instance
func (app *builder) Now() (Shader, error) {
	if app.sh == nil {
		return nil, errors.New("the shader hash is mandatory in order to build a Shader instance")
	}

	if app.id == 0 {
		return nil, errors.New("the identifier is mandatory in order to build a Shader instance")
	}

	return createShader(*app.sh, app.id), nil
}
