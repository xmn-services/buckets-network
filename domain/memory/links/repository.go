package links

import (
	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	transfer_link "github.com/xmn-services/buckets-network/domain/transfers/links"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	builder         Builder
	blockRepository mined_blocks.Repository
	trRepository    transfer_link.Repository
}

func createRepository(
	builder Builder,
	blockRepository mined_blocks.Repository,
	trRepository transfer_link.Repository,
) Repository {
	out := repository{
		builder:         builder,
		blockRepository: blockRepository,
		trRepository:    trRepository,
	}

	return &out
}

// Retrieve retrieves a link by hash
func (app *repository) Retrieve(hsh hash.Hash) (Link, error) {
	trLink, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	nextHash := trLink.Next()
	next, err := app.blockRepository.Retrieve(nextHash)
	if err != nil {
		return nil, err
	}

	prev := trLink.Previous()
	index := trLink.Index()
	createdOn := trLink.CreatedOn()
	return app.builder.Create().
		WithNext(next).
		WithPrevious(prev).
		WithIndex(index).
		CreatedOn(createdOn).
		Now()
}
