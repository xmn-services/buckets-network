package buckets

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

// Create creates a new factory instance
func (app *factory) Create() (Buckets, error) {
	now := time.Now().UTC()
	return app.builder.Create().WithoutHash().CreatedOn(now).LastUpdatedOn(now).Now()
}
