package buckets

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

// Create creates a new factory instance
func (app *factory) Create() (Buckets, error) {
	return app.builder.Create().Now()
}
