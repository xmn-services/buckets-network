package worlds

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

// Create creates a new World instance
func (app *factory) Create() (World, error) {
	return app.builder.Create().Now()
}
