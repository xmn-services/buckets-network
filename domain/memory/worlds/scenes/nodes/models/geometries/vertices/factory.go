package vertices

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

// Create creates a new Vertices instance
func (app *factory) Create() (Vertices, error) {
	return app.builder.Create().WithoutHash().Now()
}
