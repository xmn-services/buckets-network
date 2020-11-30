package contents

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	transfer_content "github.com/xmn-services/buckets-network/domain/transfers/contents"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type service struct {
	hashAdapter      hash.Adapter
	trContentService transfer_content.Service
}

func createService(
	hashAdapter hash.Adapter,
	trContentService transfer_content.Service,
) Service {
	out := service{
		hashAdapter:      hashAdapter,
		trContentService: trContentService,
	}

	return &out
}

// Save saves a chunk in bucket
func (app *service) Save(bucket buckets.Bucket, data []byte) error {
	hsh, err := app.hashAdapter.FromBytes(data)
	if err != nil {
		return err
	}

	file, _, err := bucket.FileChunkByHash(*hsh)
	if err != nil {
		return err
	}

	return app.trContentService.Save(bucket.Hash(), file.Hash(), data)
}

// Delete deletes a chunk from bucket
func (app *service) Delete(bucket buckets.Bucket, chunkHash hash.Hash) error {
	file, chunk, err := bucket.FileChunkByHash(chunkHash)
	if err != nil {
		return err
	}

	return app.trContentService.Delete(bucket.Hash(), file.Hash(), chunk.Hash())
}

// DeleteAll deletes all chunks from bucket
func (app *service) DeleteAll(bucket buckets.Bucket) error {
	return app.trContentService.DeleteAll(bucket.Hash())
}
