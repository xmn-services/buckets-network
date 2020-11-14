package identities

type application struct {
	current Current
	sub     SubApplications
}

func createApplication(
	current Current,
	sub SubApplications,
) Application {
	out := application{
		current: current,
		sub:     sub,
	}

	return &out
}

// Current returns the current application
func (app *application) Current() Current {
	return app.current
}

// Sub returns the sub application
func (app *application) Sub() SubApplications {
	return app.sub
}
