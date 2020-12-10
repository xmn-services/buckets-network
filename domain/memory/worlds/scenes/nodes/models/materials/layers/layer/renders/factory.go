package renders

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

// Create creates a new Renders instance
func (app *factory) Create() (Renders, error) {
	return app.builder.Create().WithoutHash().Now()
}
