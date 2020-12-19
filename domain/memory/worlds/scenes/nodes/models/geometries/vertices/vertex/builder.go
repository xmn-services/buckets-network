package vertex

import (
	"errors"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
)

type builder struct {
	pos *fl32.Vec3
	tex *fl32.Vec2
}

func createBuilder() Builder {
	out := builder{
		pos: nil,
		tex: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithPosition adds a position to the builder
func (app *builder) WithPosition(pos fl32.Vec3) Builder {
	app.pos = &pos
	return app
}

// WithTexture adds a texture to the builder
func (app *builder) WithTexture(tex fl32.Vec2) Builder {
	app.tex = &tex
	return app
}

// Now builds a new Vertex instance
func (app *builder) Now() (Vertex, error) {
	if app.pos == nil {
		return nil, errors.New("the position is mandatory in order to build a Vertex instance")
	}

	if app.tex == nil {
		return nil, errors.New("the texture coordinates are mandatory in order to build a Vertex instance")
	}

	return createVertex(*app.pos, *app.tex), nil
}
