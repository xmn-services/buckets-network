package identities

import (
	"github.com/xmn-services/buckets-network/application/identities/buckets"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
	"github.com/xmn-services/buckets-network/application/identities/storages"
)

type subApplications struct {
	bucket  buckets.Application
	daemon  daemons.Application
	storage storages.Application
}

func createSubApplications(
	bucket buckets.Application,
	daemon daemons.Application,
	storage storages.Application,
) SubApplications {
	out := subApplications{
		bucket:  bucket,
		daemon:  daemon,
		storage: storage,
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

// Storage returns the storage application
func (app *subApplications) Storage() storages.Application {
	return app.storage
}
