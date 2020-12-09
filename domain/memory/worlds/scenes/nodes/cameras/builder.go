package cameras

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	viewport         rectangles.Rectangle
	fov              *float64
	index            uint
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		viewport:         nil,
		fov:              nil,
		index:            0,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.immutableBuilder,
	)
}

// WithViewport adds a viewport to the builder
func (app *builder) WithViewport(viewport rectangles.Rectangle) Builder {
	app.viewport = viewport
	return app
}

// WithFieldOfView adds a fov to the builder
func (app *builder) WithFieldOfView(fov float64) Builder {
	app.fov = &fov
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.index = index
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Camera instance
func (app *builder) Now() (Camera, error) {
	if app.viewport == nil {
		return nil, errors.New("the viewport is mandatory in order to build a Camera instance")
	}

	if app.fov == nil {
		return nil, errors.New("the fieldOfView is mandatory in order to build a Camera instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.viewport.String()),
		[]byte(strconv.FormatFloat(*app.fov, 'f', 10, 64)),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createCamera(immutable, app.viewport, *app.fov, app.index), nil
}
