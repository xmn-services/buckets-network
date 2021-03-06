package buckets

import (
	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	adapter        Adapter
	fileRepository file.Repository
}

func createRepository(adapter Adapter, fileRepository file.Repository) Repository {
	out := repository{
		adapter:        adapter,
		fileRepository: fileRepository,
	}

	return &out
}

// Retrieve retrieves a bucket by hash
func (app *repository) Retrieve(hsh hash.Hash) (Bucket, error) {
	js, err := app.fileRepository.Retrieve(hsh.String())
	if err != nil {
		return nil, err
	}

	return app.adapter.ToBucket(js)
}
