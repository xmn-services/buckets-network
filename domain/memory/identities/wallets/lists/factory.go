package lists

type factory struct {
	builder Builder
}

func createFactory(builder Builder) Factory {
	out := factory{
		builder: builder,
	}

	return &out
}

// Create creates a new lists instance
func (app *factory) Create() (Lists, error) {
	return app.builder.Create().Now()
}
