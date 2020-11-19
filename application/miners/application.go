package miners

import (
	"errors"
	"fmt"
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

		// save the genesis:
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

// Init mines the initial chain
func (app *application) Init(miningValue uint8, baseDifficulty uint, increasePerBucket float64, linkDifficulty uint, rootAdditionalBuckets uint, headAdditionalBuckets uint) (chains.Chain, error) {
	_, err := app.genesisRepository.Retrieve()
	if err == nil {
		return nil, errors.New("the genesis block has already been created")
	}

	// create the genesis:
	createdOn := time.Now().UTC()
	gen, err := app.genesisBuilder.Create().
		WithMiningValue(miningValue).
		WithBlockDifficultyBase(baseDifficulty).
		WithBlockDifficultyIncreasePerBucket(increasePerBucket).
		WithLinkDifficulty(linkDifficulty).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return nil, err
	}

	// mine the root block:
	root, err := app.block(gen, []string{}, baseDifficulty, increasePerBucket, rootAdditionalBuckets)
	if err != nil {
		return nil, err
	}

	// mine the head block:
	headBlock, err := app.block(gen, []string{}, baseDifficulty, increasePerBucket, headAdditionalBuckets)
	if err != nil {
		return nil, err
	}

	// mine the head link:
	head, err := app.mineLink(root, headBlock, linkDifficulty)
	if err != nil {
		return nil, err
	}

	// build the chain:
	rootBlock := root.Block()
	rootBucketAmount := rootBlock.Additional()
	if rootBlock.HasBuckets() {
		rootBucketAmount += uint(len(rootBlock.Buckets()))
	}

	headBlockBlock := headBlock.Block()
	headBlockAmount := headBlockBlock.Additional()
	if headBlockBlock.HasBuckets() {
		headBlockAmount += uint(len(headBlockBlock.Buckets()))
	}

	totalAmount := rootBucketAmount + headBlockAmount
	return app.chainBuilder.Create().WithGenesis(gen).WithRoot(root).WithHead(head).WithTotal(totalAmount).Now()
}

// Block mines a block
func (app *application) Block(bucketHashesStr []string, baseDifficulty uint, increasePerBucket float64, additionalBuckets uint) (mined_block.Block, error) {
	// retrieve the genesis:
	gen, err := app.genesisRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	return app.block(gen, bucketHashesStr, baseDifficulty, increasePerBucket, additionalBuckets)
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

func (app *application) block(
	gen genesis.Genesis,
	bucketHashesStr []string,
	baseDifficulty uint,
	increasePerBucket float64,
	additionalBuckets uint,
) (mined_block.Block, error) {
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

	return app.mineBlock(gen, buckets, baseDifficulty, increasePerBucket, additionalBuckets)
}

func (app *application) mineBlock(
	gen genesis.Genesis,
	buckets []buckets.Bucket,
	baseDifficulty uint,
	increasePerBucket float64,
	additionalBuckets uint,
) (mined_block.Block, error) {
	// calculate the difficulty:
	amountBuckets := uint(len(buckets)) + additionalBuckets
	difficulty := blockDifficulty(
		baseDifficulty,
		increasePerBucket,
		amountBuckets,
	)

	if difficulty > maxDifficulty {
		str := fmt.Sprintf("the block cannot be mined because the required difficulty (%d) to mine it is higher than the maximum difficulty (%d), try to reduce the amount of buckets (%d) in your block in order to reduce its difficulty", difficulty, maxDifficulty, amountBuckets)
		return nil, errors.New(str)
	}

	// build the block:
	createdOn := time.Now().UTC()
	block, err := app.blockBuilder.Create().
		WithGenesis(gen).
		WithBuckets(buckets).
		WithAdditional(additionalBuckets).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return nil, err
	}

	// mine the block:
	miningValue := gen.MiningValue()
	minedCreatedOn := time.Now().UTC()
	results, err := mine(app.hashAdapter, miningValue, difficulty, block.Hash())
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
	if linkDifficulty > maxDifficulty {
		str := fmt.Sprintf("the requested link mining cannot be executed because the requested difficulty (%d) is higher than the maximum difficulty (%d)", maxDifficulty, linkDifficulty)
		return nil, errors.New(str)
	}

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
	miningValue := prevMinedBlock.Block().Genesis().MiningValue()
	results, err := mine(app.hashAdapter, miningValue, linkDifficulty, link.Hash())
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
