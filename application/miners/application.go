package miners

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter          hash.Adapter
	bucketRepository     buckets.Repository
	blockBuilder         blocks.Builder
	minedBlockBuilder    mined_block.Builder
	minedBlockRepository mined_block.Repository
	linkBuilder          links.Builder
	minedLinkBuilder     mined_link.Builder
	genesisBuilder       genesis.Builder
	genesisRepository    genesis.Repository
	chainBuilder         chains.Builder
}

func createApplication(
	hashAdapter hash.Adapter,
	bucketRepository buckets.Repository,
	blockBuilder blocks.Builder,
	minedBlockBuilder mined_block.Builder,
	minedBlockRepository mined_block.Repository,
	linkBuilder links.Builder,
	minedLinkBuilder mined_link.Builder,
	genesisBuilder genesis.Builder,
	genesisRepository genesis.Repository,
	chainBuilder chains.Builder,
) Application {
	out := application{
		hashAdapter:          hashAdapter,
		bucketRepository:     bucketRepository,
		blockBuilder:         blockBuilder,
		minedBlockBuilder:    minedBlockBuilder,
		minedBlockRepository: minedBlockRepository,
		linkBuilder:          linkBuilder,
		minedLinkBuilder:     minedLinkBuilder,
		genesisBuilder:       genesisBuilder,
		genesisRepository:    genesisRepository,
		chainBuilder:         chainBuilder,
	}

	return &out
}

// Genesis mines the genesis chain
func (app *application) Genesis(baseDifficulty uint, increasePerBucket float64, linkDifficulty uint) (chains.Chain, error) {
	_, err := app.genesisRepository.Retrieve()
	if err == nil {
		return nil, errors.New("the genesis block has already been created")
	}

	// create the genesis:
	createdOn := time.Now().UTC()
	gen, err := app.genesisBuilder.Create().
		WithBlockDifficultyBase(baseDifficulty).
		WithBlockDifficultyIncreasePerBucket(increasePerBucket).
		WithLinkDifficulty(linkDifficulty).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return nil, err
	}

	// mine the root block:
	root, err := app.Block([]string{}, baseDifficulty, increasePerBucket)
	if err != nil {
		return nil, err
	}

	// mine the head block:
	headBlock, err := app.Block([]string{}, baseDifficulty, increasePerBucket)
	if err != nil {
		return nil, err
	}

	// mine the head link:
	head, err := app.mineLink(root, headBlock, linkDifficulty)
	if err != nil {
		return nil, err
	}

	// build the chain:
	return app.chainBuilder.Create().WithGenesis(gen).WithRoot(root).WithHead(head).WithTotal(0).Now()
}

// Block mines a block
func (app *application) Block(bucketHashesStr []string, baseDifficulty uint, increasePerBucket float64) (mined_block.Block, error) {
	bucketHashes := []hash.Hash{}
	for _, oneBucketHashStr := range bucketHashesStr {
		bucketHash, err := app.hashAdapter.FromString(oneBucketHashStr)
		if err != nil {
			return nil, err
		}

		bucketHashes = append(bucketHashes, *bucketHash)
	}

	buckets, err := app.bucketRepository.RetrieveAll(bucketHashes)
	if err != nil {
		return nil, err
	}

	return app.mineBlock(buckets, baseDifficulty, increasePerBucket)
}

// Link mines a link
func (app *application) Link(prevMinedBlockHasStr string, newMinedBlockHashStr string, linkDifficulty uint) (mined_link.Link, error) {
	prevMinedBlockHash, err := app.hashAdapter.FromString(prevMinedBlockHasStr)
	if err != nil {
		return nil, err
	}

	newMinedBlockHash, err := app.hashAdapter.FromString(newMinedBlockHashStr)
	if err != nil {
		return nil, err
	}

	prevMinedBlock, err := app.minedBlockRepository.Retrieve(*prevMinedBlockHash)
	if err != nil {
		return nil, err
	}

	newMinedBlock, err := app.minedBlockRepository.Retrieve(*newMinedBlockHash)
	if err != nil {
		return nil, err
	}

	return app.mineLink(prevMinedBlock, newMinedBlock, linkDifficulty)
}

func (app *application) mineBlock(buckets []buckets.Bucket, baseDifficulty uint, increasePerBucket float64) (mined_block.Block, error) {
	// calculate the difficulty:
	difficulty := blockDifficulty(
		baseDifficulty,
		increasePerBucket,
		uint(len(buckets)),
	)

	// retrieve the genesis:
	gen, err := app.genesisRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	// build the block:
	createdOn := time.Now().UTC()
	block, err := app.blockBuilder.Create().
		WithGenesis(gen).
		WithBuckets(buckets).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return nil, err
	}

	// mine the block:
	minedCreatedOn := time.Now().UTC()
	results, err := mine(app.hashAdapter, difficulty, block.Hash())
	if err != nil {
		return nil, err
	}

	return app.minedBlockBuilder.Create().
		WithBlock(block).
		WithMining(results).
		CreatedOn(minedCreatedOn).
		Now()
}

func (app *application) mineLink(prevMinedBlock mined_block.Block, newMinedBlock mined_block.Block, linkDifficulty uint) (mined_link.Link, error) {
	prev := prevMinedBlock.Hash()
	linkCreatedOn := time.Now().UTC()
	link, err := app.linkBuilder.Create().
		WithPreviousLink(prev).
		WithNext(newMinedBlock).
		CreatedOn(linkCreatedOn).
		Now()

	if err != nil {
		return nil, err
	}

	// mine:
	results, err := mine(app.hashAdapter, linkDifficulty, link.Hash())
	if err != nil {
		return nil, err
	}

	// return the mined link:
	minedLinkCreatedOn := time.Now().UTC()
	return app.minedLinkBuilder.Create().
		WithLink(link).
		WithMining(results).
		CreatedOn(minedLinkCreatedOn).
		Now()
}
