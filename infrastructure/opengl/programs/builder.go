package programs

import (
	"errors"

	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/program"
)

type builder struct {
	list []program.Program
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

// WithPrograms add programs to the builder
func (app *builder) WithPrograms(progs []program.Program) Builder {
	app.list = progs
	return app
}

// Now builds a new Programs instance
func (app *builder) Now() (Programs, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("the compiled programs are mandatory in order to build a Programs instance")
	}

	return createPrograms(app.list), nil
}
