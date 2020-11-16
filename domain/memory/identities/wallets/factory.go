package wallets

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

// Create creates a new Wallet instance
func (app *factory) Create() (Wallet, error) {
	return app.builder.Create().Now()
}
