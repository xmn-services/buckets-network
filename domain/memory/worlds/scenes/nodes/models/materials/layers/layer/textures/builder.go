package textures

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/rows"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	dimension        rectangles.Rectangle
	pixels           rows.Rows
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		dimension:        nil,
		pixels:           nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithDimension adds a dimension to the builder
func (app *builder) WithDimension(dimension rectangles.Rectangle) Builder {
	app.dimension = dimension
	return app
}

// WithPixels adds pixels to the builder
func (app *builder) WithPixels(pixels rows.Rows) Builder {
	app.pixels = pixels
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Texture instance
func (app *builder) Now() (Texture, error) {
	if app.dimension == nil {
		return nil, errors.New("the dimension is mandatory in order to build a Texture instance")
	}

	if app.pixels == nil {
		return nil, errors.New("the pixels are mandatory in order to build a Texture instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.dimension.String()),
		app.pixels.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createTexture(immutable, app.dimension, app.pixels), nil
}
