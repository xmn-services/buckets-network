package layers

type factory struct {
	builder Builder
}

func createFactory(
	builder Builder,
) Factory {
	out := factory{
		builder: builder,
	}

	return &out
}

// Create creates a new Layers instance
func (app *factory) Create() (Layers, error) {
	return app.builder.Create().WithoutHash().Now()
}
