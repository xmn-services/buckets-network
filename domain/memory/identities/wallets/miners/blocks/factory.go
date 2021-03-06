package blocks

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

// Create generates a new blocks instance
func (app *factory) Create() (Blocks, error) {
	return app.builder.Create().Now()
}
