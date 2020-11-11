package blocks

import "time"

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
	now := time.Now().UTC()
	return app.builder.Create().WithoutHash().CreatedOn(now).LastUpdatedOn(now).Now()
}
