package worlds

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	sceneFactory     scenes.Factory
	scenes           []scenes.Scene
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
	sceneFactory scenes.Factory,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		sceneFactory:     sceneFactory,
		scenes:           nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder, app.sceneFactory)
}

// WithScenes add scenes to the builder
func (app *builder) WithScenes(scenes []scenes.Scene) Builder {
	app.scenes = scenes
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new World instance
func (app *builder) Now() (World, error) {
	if app.scenes != nil && len(app.scenes) <= 0 {
		app.scenes = nil
	}

	if app.scenes == nil {
		scene, err := app.sceneFactory.Create()
		if err != nil {
			return nil, err
		}

		app.scenes = []scenes.Scene{
			scene,
		}
	}

	data := [][]byte{}
	for _, oneScene := range app.scenes {
		data = append(data, oneScene.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createWorld(immutable, app.scenes), nil
}
