package rows

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

// Create creates a new Rows instance
func (app *factory) Create() (Rows, error) {
	return app.builder.Create().WithoutHash().Now()
}
