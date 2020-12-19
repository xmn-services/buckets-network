package spaces

import (
	"errors"

	"github.com/go-gl/mathgl/mgl32"
)

type builder struct {
	initial     Space
	pos         *mgl32.Vec3
	orientation *mgl32.Vec4
}

func createBuilder() Builder {
	out := builder{
		initial:     nil,
		pos:         nil,
		orientation: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInitial adds an initial space to the builder
func (app *builder) WithInitial(space Space) Builder {
	app.initial = space
	return app
}

// WithPosition adds a position to the builder
func (app *builder) WithPosition(pos mgl32.Vec3) Builder {
	app.pos = &pos
	return app
}

// WithOrientation adds an orientation to the builder
func (app *builder) WithOrientation(orientation mgl32.Vec4) Builder {
	app.orientation = &orientation
	return app
}

// Now builds a new Space instance
func (app *builder) Now() (Space, error) {
	if app.pos == nil {
		return nil, errors.New("the position is mandatory in order to build a Space instance")
	}

	if app.orientation == nil {
		return nil, errors.New("the orientation is mandatory in order to build a Space instance")
	}

	space := createSpace(*app.pos, *app.orientation)
	if app.initial == nil {
		return space, nil
	}

	return app.initial.Add(space), nil
}
