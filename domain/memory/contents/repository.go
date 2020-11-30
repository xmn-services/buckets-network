package contents

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	transfer_content "github.com/xmn-services/buckets-network/domain/transfers/contents"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	trContentRepository transfer_content.Repository
}

func createRepository(
	trContentRepository transfer_content.Repository,
) Repository {
	out := repository{
		trContentRepository: trContentRepository,
	}

	return &out
}

// Retrieve retrieves a chunk's data
func (app *repository) Retrieve(bucket buckets.Bucket, fileHash hash.Hash, chunkHash hash.Hash) ([]byte, error) {
	return app.trContentRepository.Retrieve(bucket.Hash(), fileHash, chunkHash)
}
