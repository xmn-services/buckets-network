package chains

import (
	"errors"
	"fmt"

	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	transfer_chains "github.com/xmn-services/buckets-network/domain/transfers/chains"
)

type repository struct {
	genesisRepository genesis.Repository
	blockRepository   mined_block.Repository
	linkRepository    mined_link.Repository
	trRepository      transfer_chains.Repository
	builder           Builder
}

func createRepository(
	genesisRepository genesis.Repository,
	blockRepository mined_block.Repository,
	linkRepository mined_link.Repository,
	trRepository transfer_chains.Repository,
	builder Builder,
) Repository {
	out := repository{
		genesisRepository: genesisRepository,
		blockRepository:   blockRepository,
		linkRepository:    linkRepository,
		trRepository:      trRepository,
		builder:           builder,
	}

	return &out
}

// RetrieveAtIndex retrieves chain at index
func (app *repository) RetrieveAtIndex(index uint) (Chain, error) {
	trChain, err := app.trRepository.RetrieveAtIndex(index)
	if err != nil {
		return nil, err
	}

	return app.toChain(trChain)
}

// Retrieve retrieves a chain instance
func (app *repository) Retrieve() (Chain, error) {
	trChain, err := app.trRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	return app.toChain(trChain)
}

func (app *repository) toChain(trChain transfer_chains.Chain) (Chain, error) {
	genHash := trChain.Genesis()
	gen, err := app.genesisRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	if !genHash.Compare(gen.Hash()) {
		str := fmt.Sprintf("the stored genesis hash, in the chain's stored data is invalid (expected: %s, returned: %s)", genHash.String(), gen.Hash().String())
		return nil, errors.New(str)
	}

	rootHash := trChain.Root()
	root, err := app.blockRepository.Retrieve(rootHash)
	if err != nil {
		return nil, err
	}

	headHash := trChain.Head()
	head, err := app.linkRepository.Retrieve(headHash)
	if err != nil {
		return nil, err
	}

	total := trChain.Total()
	createdOn := trChain.CreatedOn()
	return app.builder.Create().
		WithGenesis(gen).
		WithRoot(root).
		WithHead(head).
		WithTotal(total).
		CreatedOn(createdOn).
		Now()
}
