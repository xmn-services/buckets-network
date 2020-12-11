package layer

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	opacity          float64
	viewport         rectangles.Rectangle
	renders          renders.Renders
	shaders          shaders.Shaders
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		opacity:          float64(1),
		viewport:         nil,
		renders:          nil,
		shaders:          nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithOpacity adds an opacity to the builder
func (app *builder) WithOpacity(opacity float64) Builder {
	app.opacity = opacity
	return app
}

// WithViewport adds a viewport to the builder
func (app *builder) WithViewport(viewport rectangles.Rectangle) Builder {
	app.viewport = viewport
	return app
}

// WithRenders add renders to the builder
func (app *builder) WithRenders(renders renders.Renders) Builder {
	app.renders = renders
	return app
}

// WithShaders add shaders to the builder
func (app *builder) WithShaders(shaders shaders.Shaders) Builder {
	app.shaders = shaders
	return app
}

// CreatedOn add a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Layer instance
func (app *builder) Now() (Layer, error) {
	if app.viewport == nil {
		return nil, errors.New("the viewport is mandatory in order to build a Layer instance")
	}

	if app.renders == nil {
		return nil, errors.New("the renders is mandatory in order to build a Layer instance")
	}

	if app.shaders == nil {
		return nil, errors.New("the shaders is mandatory in order to build a Layer instance")
	}

	if app.opacity < 0.0 || app.opacity > 1.0 {
		return nil, errors.New("the opacity must be a float between 0.0 and 1.0")
	}

	if !app.shaders.IsFragment() {
		return nil, errors.New("the material's layer shaders were expected to be fragment shaders")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.FormatFloat(app.opacity, 'f', 10, 64)),
		[]byte(app.viewport.String()),
		app.renders.Hash().Bytes(),
		app.shaders.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createLayer(immutable, app.opacity, app.viewport, app.renders, app.shaders), nil
}
