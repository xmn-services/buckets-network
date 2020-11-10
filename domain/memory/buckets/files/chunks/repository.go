package chunks

import (
	transfer_chunk "github.com/xmn-services/buckets-network/domain/transfers/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	trRepository transfer_chunk.Repository
	builder      Builder
}

func createRepository(
	trRepository transfer_chunk.Repository,
	builder Builder,
) Repository {
	out := repository{
		trRepository: trRepository,
		builder:      builder,
	}

	return &out
}

// Retrieve retrieves a chunk instance by hash
func (app *repository) Retrieve(hash hash.Hash) (Chunk, error) {
	trChunk, err := app.trRepository.Retrieve(hash)
	if err != nil {
		return nil, err
	}

	sizeInBytes := trChunk.SizeInBytes()
	data := trChunk.Data()
	createdOn := trChunk.CreatedOn()
	return app.builder.Create().WithSizeInBytes(sizeInBytes).WithData(data).CreatedOn(createdOn).Now()
}

// RetrieveAll retrieves all chunk instances by hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]Chunk, error) {
	out := []Chunk{}
	for _, oneHash := range hashes {
		chk, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, chk)
	}

	return out, nil
}
