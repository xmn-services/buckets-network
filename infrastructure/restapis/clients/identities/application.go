package identities

import "github.com/xmn-services/buckets-network/application/commands/identities"

type application struct {
	current identities.Current
	sub     identities.SubApplications
}

func createApplication(
	current identities.Current,
	sub identities.SubApplications,
) identities.Application {
	out := application{
		current: current,
		sub:     sub,
	}

	return &out
}

// Current returns the current application
func (obj *application) Current() identities.Current {
	return obj.current
}

// Sub returns the subApplications
func (obj *application) Sub() identities.SubApplications {
	return obj.sub
}
