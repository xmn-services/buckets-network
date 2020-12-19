package scenes

import (
	"errors"

	domain_scenes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/nodes"
)

type builder struct {
	nodesBuilder       nodes.Builder
	defaultCameraIndex uint
	cameraIndex        uint
	scene              domain_scenes.Scene
}

func createBuilder(
	nodesBuilder nodes.Builder,
	defaultCameraIndex uint,
) Builder {
	out := builder{
		nodesBuilder:       nodesBuilder,
		defaultCameraIndex: defaultCameraIndex,
		cameraIndex:        defaultCameraIndex,
		scene:              nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.nodesBuilder, app.defaultCameraIndex)
}

// WithScene adds a scene to the builder
func (app *builder) WithScene(scene domain_scenes.Scene) Builder {
	app.scene = scene
	return app
}

// WithSceWithCameraIndexne adds a cameraIndex to the builder
func (app *builder) WithCameraIndex(cameraIndex uint) Builder {
	app.cameraIndex = cameraIndex
	return app
}

// Now builds a new Scene instance
func (app *builder) Now() (Scene, error) {
	if app.scene == nil {
		return nil, errors.New("the scene is mandatory in order to build a Scene instance")
	}

	if app.scene.HasNodes() {
		domainNodes := app.scene.Nodes()
		list, err := app.nodesBuilder.Create().WithNodes(domainNodes).Now()
		if err != nil {
			return nil, err
		}

		return createSceneWithNodes(app.scene, app.cameraIndex, list), nil
	}

	return createScene(app.scene, app.cameraIndex), nil
}
