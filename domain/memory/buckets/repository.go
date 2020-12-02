package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	transfer_bucket "github.com/xmn-services/buckets-network/domain/transfers/buckets"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	fileRepository files.Repository
	trRepository   transfer_bucket.Repository
	builder        Builder
}

func createRepository(
	fileRepository files.Repository,
	trRepository transfer_bucket.Repository,
	builder Builder,
) Repository {
	out := repository{
		fileRepository: fileRepository,
		trRepository:   trRepository,
		builder:        builder,
	}

	return &out
}

// Retrieve retrieves an bucket instance by hash
func (app *repository) Retrieve(hsh hash.Hash) (Bucket, error) {
	trBucket, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	amount := trBucket.Amount()
	fileHashes := []hash.Hash{}
	leaves := trBucket.Files().Parent().BlockLeaves().Leaves()
	for i := 0; i < int(amount); i++ {
		fileHashes = append(fileHashes, leaves[i].Head())
	}

	files, err := app.fileRepository.RetrieveAll(fileHashes)
	if err != nil {
		return nil, err
	}

	createdOn := trBucket.CreatedOn()
	return app.builder.Create().WithFiles(files).CreatedOn(createdOn).Now()
}

// RetrieveAll retrieves buckets from hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]Bucket, error) {
	out := []Bucket{}
	for _, oneHash := range hashes {
		bucket, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, bucket)
	}

	return out, nil
}

// RetrieveWithChunksContentFromPath retrieves a bucket with chunk's contents from path
func (app *repository) RetrieveWithChunksContentFromPath(path string, decryptPubKey public.Key) (Bucket, [][][]byte, error) {
	files, chunksContent, err := app.fileRepository.RetrieveAllWithChunksContentFromPath(path, decryptPubKey)
	if err != nil {
		return nil, nil, err
	}

	bucket, err := app.builder.Create().WithFiles(files).Now()
	if err != nil {
		return nil, nil, err
	}

	return bucket, chunksContent, nil
}
