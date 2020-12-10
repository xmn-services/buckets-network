package pixels

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

// Create creates a new Pixels instance
func (app *factory) Create() (Pixels, error) {
	return app.builder.Create().WithoutHash().Now()
}
