package nodes

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter          hash.Adapter
	immutableBuilder     entities.ImmutableBuilder
	position             *fl32.Vec3
	orientationAngle     *float32
	orientationDirection *fl32.Vec3
	model                models.Model
	camera               cameras.Camera
	children             []Node
	createdOn            *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:          hashAdapter,
		immutableBuilder:     immutableBuilder,
		position:             nil,
		orientationAngle:     nil,
		orientationDirection: nil,
		model:                nil,
		camera:               nil,
		children:             nil,
		createdOn:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithPosition adds a position to the builder
func (app *builder) WithPosition(pos fl32.Vec3) Builder {
	app.position = &pos
	return app
}

// WithOrientationAngle adds an orientation angle to the builder
func (app *builder) WithOrientationAngle(angle float32) Builder {
	app.orientationAngle = &angle
	return app
}

// WithOrientationDirection adds an orientation direction to the builder
func (app *builder) WithOrientationDirection(direction fl32.Vec3) Builder {
	app.orientationDirection = &direction
	return app
}

// WithModel adds a model to the builder
func (app *builder) WithModel(model models.Model) Builder {
	app.model = model
	return app
}

// WithCamera adds a camera to the builder
func (app *builder) WithCamera(camera cameras.Camera) Builder {
	app.camera = camera
	return app
}

// WithChildren adds children nodes to the builder
func (app *builder) WithChildren(children []Node) Builder {
	app.children = children
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Node instance
func (app *builder) Now() (Node, error) {
	if app.position == nil {
		return nil, errors.New("the position is mandatory in order to build a Node instance")
	}

	if app.orientationAngle == nil {
		return nil, errors.New("the orientation angle is mandatory in order to build a Node instance")
	}

	if app.orientationDirection == nil {
		return nil, errors.New("the orientation direction is mandatory in order to build a Node instance")
	}

	var content Content
	if app.model != nil {
		content = createContentWithModel(app.model)
	}

	if app.camera != nil {
		content = createContentWithCamera(app.camera)
	}

	if content == nil {
		return nil, errors.New("the content (model or camera) is mandatory in order to build a Node instance")
	}

	if app.children != nil && len(app.children) <= 0 {
		app.children = nil
	}

	data := [][]byte{
		content.Hash().Bytes(),
		[]byte(app.position.String()),
		[]byte(app.orientationDirection.String()),
		[]byte(strconv.FormatFloat(float64(*app.orientationAngle), 'f', 10, 32)),
	}

	if app.children != nil {
		for _, oneNode := range app.children {
			data = append(data, oneNode.Hash().Bytes())
		}
	}

	if app.model != nil {
		data = append(data, app.model.Hash().Bytes())
	}

	if app.camera != nil {
		data = append(data, app.camera.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	orientation := createOrientation(*app.orientationAngle, *app.orientationDirection)
	if content != nil && app.children != nil {
		return createNodeWithContentAndNodes(immutable, *app.position, orientation, content, app.children), nil
	}

	if content != nil {
		return createNodeWithContent(immutable, *app.position, orientation, content), nil
	}

	if app.children != nil {
		return createNodeWithNodes(immutable, *app.position, orientation, app.children), nil
	}

	return createNode(immutable, *app.position, orientation), nil
}
