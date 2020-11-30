package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter       hash.Adapter
	bucketRepository  buckets.Repository
	contentRepository contents.Repository
}

func createApplication(
	hashAdapter hash.Adapter,
	bucketRepository buckets.Repository,
	contentRepository contents.Repository,
) Application {
	out := application{
		hashAdapter:       hashAdapter,
		bucketRepository:  bucketRepository,
		contentRepository: contentRepository,
	}

	return &out
}

// Exists returns true if the chunk exists, false otherwise
func (app *application) Exists(bucketHashStr string, chunkHashStr string) bool {
	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return false
	}

	chunkHash, err := app.hashAdapter.FromString(chunkHashStr)
	if err != nil {
		return false
	}

	bucket, err := app.bucketRepository.Retrieve(*bucketHash)
	if err != nil {
		return false
	}

	_, _, err = bucket.FileChunkByHash(*chunkHash)
	if err != nil {
		return false
	}

	return true
}

// Retrieve retrieves a chunk's data
func (app *application) Retrieve(bucketHashStr string, chunkHashStr string) ([]byte, error) {
	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return nil, err
	}

	chunkHash, err := app.hashAdapter.FromString(chunkHashStr)
	if err != nil {
		return nil, err
	}

	bucket, err := app.bucketRepository.Retrieve(*bucketHash)
	if err != nil {
		return nil, err
	}

	file, chunk, err := bucket.FileChunkByHash(*chunkHash)
	if err != nil {
		return nil, err
	}

	return app.contentRepository.Retrieve(bucket, file.Hash(), chunk.Hash())
}
