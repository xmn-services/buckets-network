package miners

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

// Create creates a new miner instance
func (app *factory) Create() (Miner, error) {
	return app.builder.Create().Now()
}
