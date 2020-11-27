package buckets

import (
	command_bucket "github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
)

type application struct {
}

func createApplication() command_bucket.Application {
	out := application{}
	return &out
}

// Add adds a path to the bucket application
func (app *application) Add(absolutePath string) error {
	return nil
}

// Delete deletes a bucket by hash
func (app *application) Delete(hashStr string) error {
	return nil
}

// Retrieve retrieves a bucket by hash
func (app *application) Retrieve(hashStr string) (buckets.Bucket, error) {
	return nil, nil
}

// RetrieveAll retrieves all buckets
func (app *application) RetrieveAll() ([]buckets.Bucket, error) {
	return nil, nil
}
