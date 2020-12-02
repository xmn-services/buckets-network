package contacts

type factory struct {
	builder Builder
}

func createFactory(builder Builder) Factory {
	out := factory{
		builder: builder,
	}

	return &out
}

// Create creates a new contacts instance
func (app *factory) Create() (Contacts, error) {
	return app.builder.Create().Now()
}
