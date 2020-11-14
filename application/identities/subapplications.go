package identities

import (
	"github.com/xmn-services/buckets-network/application/identities/buckets"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
)

type subApplications struct {
	bucket buckets.Application
	daemon daemons.Application
}

func createSubApplications(
	bucket buckets.Application,
	daemon daemons.Application,
) SubApplications {
	out := subApplications{
		bucket: bucket,
		daemon: daemon,
	}

	return &out
}

// Bucket returns the bucket application
func (app *subApplications) Bucket() buckets.Application {
	return app.bucket
}

// Daemon returns the daemon application
func (app *subApplications) Daemon() daemons.Application {
	return app.daemon
}
