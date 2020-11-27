package clients

import "github.com/xmn-services/buckets-network/application/commands"

type application struct {
	current commands.Current
	sub     commands.SubApplications
}

func createApplication(
	current commands.Current,
	sub commands.SubApplications,
) commands.Application {
	out := application{
		current: current,
		sub:     sub,
	}

	return &out
}

// Current returns the current application
func (obj *application) Current() commands.Current {
	return obj.current
}

// Sub returns the subApplication
func (obj *application) Sub() commands.SubApplications {
	return obj.sub
}
