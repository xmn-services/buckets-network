package worlds

import (
	"errors"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	domain_worlds "github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/scenes"
)

type builder struct {
	sceneBuilder scenes.Builder
	world        domain_worlds.World
	sceneIndex   uint
}

func createBuilder(
	sceneBuilder scenes.Builder,
	defaultSceneIndex uint,
) Builder {
	out := builder{
		sceneBuilder: sceneBuilder,
		world:        nil,
		sceneIndex:   defaultSceneIndex,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.sceneBuilder, app.sceneIndex)
}

// WithWorld adds a world to the builder
func (app *builder) WithWorld(world domain_worlds.World) Builder {
	app.world = world
	return app
}

// WithSceneIndex adds a sceneIndex to the builder
func (app *builder) WithSceneIndex(sceneIndex uint) Builder {
	app.sceneIndex = sceneIndex
	return app
}

//  Now builds a new World instance
func (app *builder) Now() (World, error) {
	if app.world == nil {
		return nil, errors.New("the world is mandatory in order to build a World instance")
	}

	// init OpenGL:
	err := gl.Init()
	if err != nil {
		return nil, err
	}

	// retrieve the OpenGL version:
	version := gl.GoStr(gl.GetString(gl.VERSION))
	if version == "" {
		return nil, errors.New("could not fetch OpenGL version")
	}

	// lof the OpenGL version:
	log.Println("OpenGL version", version)

	// if the world has scenes:
	if app.world.HasScenes() {
		domainScene, err := app.world.Scene(app.sceneIndex)
		if err != nil {
			return nil, err
		}

		scene, err := app.sceneBuilder.Create().WithScene(domainScene).Now()
		if err != nil {
			return nil, err
		}

		return createWorldWithScenes(app.world, app.sceneIndex, scene), nil
	}

	// if there is no scene:
	return createWorld(app.world, app.sceneIndex), nil
}
