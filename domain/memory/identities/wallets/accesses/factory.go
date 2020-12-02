package accesses

type factory struct {
	builder Builder
}

func createFactory(builder Builder) Factory {
	out := factory{
		builder: builder,
	}

	return &out
}

// Create creates a new accesses instance
func (app *factory) Create() (Accesses, error) {
	return app.builder.Create().Now()
}
