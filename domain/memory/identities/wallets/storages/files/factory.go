package files

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

// Create create files
func (app *factory) Create() (Files, error) {
	return app.builder.Create().Now()
}
