package shaders

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

// Create creates a new Shaders instance
func (app *factory) Create() (Shaders, error) {
	return app.builder.Create().WithoutHash().Now()
}
