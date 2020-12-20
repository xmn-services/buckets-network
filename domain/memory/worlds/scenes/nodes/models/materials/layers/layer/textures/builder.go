package textures

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	viewport         ints.Rectangle
	pixels           []pixels.Pixel
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
		pixels:           nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithViewport adds a viewport to the builder
func (app *builder) WithViewport(viewport ints.Rectangle) Builder {
	app.viewport = viewport
	return app
}

// WithPixels adds pixels to the builder
func (app *builder) WithPixels(pixels []pixels.Pixel) Builder {
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
	if app.viewport == nil {
		return nil, errors.New("the viewport is mandatory in order to build a Texture instance")
	}

	if app.pixels != nil && len(app.pixels) <= 0 {
		app.pixels = nil
	}

	if app.pixels == nil {
		return nil, errors.New("the pixels are mandatory in order to build a Texture instance")
	}

	position := app.viewport.Position()
	dimension := app.viewport.Dimension()
	width := (position.X() + dimension.X())
	height := (position.Y() + dimension.Y())
	expectedTotal := width * height
	pixAmount := len(app.pixels)
	if expectedTotal != pixAmount {
		str := fmt.Sprintf("the texture (width: %d, height: %d) was expecting %d pixels, %d provided", width, height, expectedTotal, pixAmount)
		return nil, errors.New(str)
	}

	data := [][]byte{
		[]byte(app.viewport.String()),
	}

	for _, onePixel := range app.pixels {
		data = append(data, []byte(onePixel.String()))
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createTexture(immutable, app.viewport, app.pixels), nil
}
