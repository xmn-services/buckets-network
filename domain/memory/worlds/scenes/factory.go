package scenes

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

// Create creates a new Scene instance
func (app *factory) Create() (Scene, error) {
	return app.builder.Create().Now()
}
