package identities

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
)

type subApplications struct {
	bucket  buckets.Application
	storage storages.Application
}

func createSubApplications(
	bucket buckets.Application,
	storage storages.Application,
) SubApplications {
	out := subApplications{
		bucket:  bucket,
		storage: storage,
	}

	return &out
}

// Bucket returns the bucket application
func (app *subApplications) Bucket() buckets.Application {
	return app.bucket
}

// Storage returns the storage application
func (app *subApplications) Storage() storages.Application {
	return app.storage
}
