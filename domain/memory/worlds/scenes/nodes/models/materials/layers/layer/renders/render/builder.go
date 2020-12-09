package render

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/colors"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	opacity          float64
	viewport         rectangles.Rectangle
	tex              textures.Texture
	camera           cameras.Camera
	color            colors.Color
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
		tex:              nil,
		camera:           nil,
		color:            nil,
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

// WithTexture adds a texture to the builder
func (app *builder) WithTexture(tex textures.Texture) Builder {
	app.tex = tex
	return app
}

// WithCamera adds a camera to the builder
func (app *builder) WithCamera(cam cameras.Camera) Builder {
	app.camera = cam
	return app
}

// WithColor adds a color to the builder
func (app *builder) WithColor(color colors.Color) Builder {
	app.color = color
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Render instanceCreatedOn adds a creation time to the builder
func (app *builder) Now() (Render, error) {
	if app.viewport == nil {
		return nil, errors.New("the viewport is mandatory in order to build a Render instance")
	}

	if app.opacity < 0.0 || app.opacity > 1.0 {
		return nil, errors.New("the opacity must be a float between 0.0 and 1.0")
	}

	var content Content
	if app.tex != nil {
		content = createContentWithTexture(app.tex)
	}

	if app.camera != nil {
		content = createContentWithCamera(app.camera)
	}

	if app.color != nil {
		content = createContentWithColor(app.color)
	}

	if content == nil {
		return nil, errors.New("the Render is  invalid")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		content.Hash().Bytes(),
		[]byte(strconv.FormatFloat(app.opacity, 'f', 10, 64)),
		[]byte(app.viewport.String()),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createRender(immutable, app.opacity, app.viewport, content), nil
}
