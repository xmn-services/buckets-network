package storages

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

// Create creates a new Storage instance
func (app *factory) Create() (Storage, error) {
	return app.builder.Create().Now()
}
