package materials

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	alpha            uint8
	viewport         ints.Rectangle
	layers           layers.Layers
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		alpha:            uint8(0),
		viewport:         nil,
		layers:           nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithAlpha adds an alpha to the builder
func (app *builder) WithAlpha(alpha uint8) Builder {
	app.alpha = alpha
	return app
}

// WithViewport adds a viewport to the builder
func (app *builder) WithViewport(viewport ints.Rectangle) Builder {
	app.viewport = viewport
	return app
}

// WithLayers add layers to the builder
func (app *builder) WithLayers(layers layers.Layers) Builder {
	app.layers = layers
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Material instance
func (app *builder) Now() (Material, error) {
	if app.viewport == nil {
		return nil, errors.New("the viewport is mandatory in order to build a Material instance")
	}

	if app.layers == nil {
		return nil, errors.New("the layers are mandatory in order to build a Material instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.alpha))),
		[]byte(app.viewport.String()),
		app.layers.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createMaterial(immutable, app.alpha, app.viewport, app.layers), nil
}
