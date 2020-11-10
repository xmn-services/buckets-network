package blocks

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	transfer_block "github.com/xmn-services/buckets-network/domain/transfers/blocks"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	builder           Builder
	genesisRepository genesis.Repository
	bucketRepository  buckets.Repository
	trRepository      transfer_block.Repository
}

func createRepository(
	builder Builder,
	genesisRepository genesis.Repository,
	bucketRepository buckets.Repository,
	trRepository transfer_block.Repository,
) Repository {
	out := repository{
		builder:           builder,
		genesisRepository: genesisRepository,
		bucketRepository:  bucketRepository,
		trRepository:      trRepository,
	}

	return &out
}

// Retrieve retrieves a block by hash
func (app *repository) Retrieve(hsh hash.Hash) (Block, error) {
	trBlock, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	gen, err := app.genesisRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	bucketHashes := []hash.Hash{}
	amountTrx := trBlock.Amount()
	leaves := trBlock.Buckets().Parent().BlockLeaves().Leaves()
	for i := 0; i < int(amountTrx); i++ {
		bucketHashes = append(bucketHashes, leaves[i].Head())
	}

	buckets, err := app.bucketRepository.RetrieveAll(bucketHashes)
	if err != nil {
		return nil, err
	}

	additional := trBlock.Additional()
	createdOn := trBlock.CreatedOn()
	return app.builder.Create().
		WithGenesis(gen).
		WithAdditional(additional).
		WithBuckets(buckets).
		CreatedOn(createdOn).
		Now()
}
