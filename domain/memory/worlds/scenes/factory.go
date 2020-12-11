package scenes

type factory struct {
	builder           Builder
	currentSceneIndex uint
}

func createFactory(
	builder Builder,
	currentSceneIndex uint,
) Factory {
	out := factory{
		builder:           builder,
		currentSceneIndex: currentSceneIndex,
	}

	return &out
}

// Create creates a new Scene instance
func (app *factory) Create() (Scene, error) {
	return app.builder.Create().WithIndex(app.currentSceneIndex).Now()
}
