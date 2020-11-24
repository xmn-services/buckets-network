package chains

import (
	"errors"

	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
)

type application struct {
	minerApplication   miners.Application
	identityRepository identities.Repository
	identityService    identities.Service
	genesisBuilder     genesis.Builder
	genesisRepository  genesis.Repository
	genesisService     genesis.Service
	blockBuilder       blocks.Builder
	blockService       blocks.Service
	minedBlockBuilder  mined_block.Builder
	linkBuilder        links.Builder
	linkService        links.Service
	minedLinkBuilder   mined_link.Builder
	chainBuilder       chains.Builder
	chainRepository    chains.Repository
	chainService       chains.Service
	name               string
	password           string
	seed               string
}

func createApplication(
	minerApplication miners.Application,
	identityRepository identities.Repository,
	identityService identities.Service,
	genesisBuilder genesis.Builder,
	genesisRepository genesis.Repository,
	genesisService genesis.Service,
	blockBuilder blocks.Builder,
	blockService blocks.Service,
	minedBlockBuilder mined_block.Builder,
	linkBuilder links.Builder,
	linkService links.Service,
	minedLinkBuilder mined_link.Builder,
	chainBuilder chains.Builder,
	chainRepository chains.Repository,
	chainService chains.Service,
	name string,
	password string,
	seed string,
) Application {
	out := application{
		minerApplication:   minerApplication,
		identityRepository: identityRepository,
		identityService:    identityService,
		genesisBuilder:     genesisBuilder,
		genesisRepository:  genesisRepository,
		genesisService:     genesisService,
		blockBuilder:       blockBuilder,
		blockService:       blockService,
		minedBlockBuilder:  minedBlockBuilder,
		linkBuilder:        linkBuilder,
		linkService:        linkService,
		minedLinkBuilder:   minedLinkBuilder,
		chainBuilder:       chainBuilder,
		chainRepository:    chainRepository,
		chainService:       chainService,
		name:               name,
		password:           password,
		seed:               seed,
	}

	return &out
}

// Init initializes the chain
func (app *application) Init(
	miningValue uint8,
	baseDifficulty uint,
	increasePerBucket float64,
	linkDifficulty uint,
	rootAdditionalBuckets uint,
	headAdditionalBuckets uint,
) error {
	_, err := app.genesisRepository.Retrieve()
	if err == nil {
		return errors.New("the genesis block has already been created")
	}

	// create the genesis:
	gen, err := app.genesisBuilder.Create().
		WithMiningValue(miningValue).
		WithBlockDifficultyBase(baseDifficulty).
		WithBlockDifficultyIncreasePerBucket(increasePerBucket).
		WithLinkDifficulty(linkDifficulty).
		Now()

	if err != nil {
		return err
	}

	// save the genesis:
	err = app.genesisService.Save(gen)
	if err != nil {
		return err
	}

	// mine the root block:
	root, err := app.mineEmptyBlock(gen, rootAdditionalBuckets)
	if err != nil {
		return err
	}

	// mine the head block:
	headBlock, err := app.mineEmptyBlock(gen, headAdditionalBuckets)
	if err != nil {
		return err
	}

	// build the head link:
	headLink, err := app.linkBuilder.Create().WithPrevious(root.Hash()).WithNext(headBlock).WithIndex(0).Now()
	if err != nil {
		return err
	}

	// save the head link:
	err = app.linkService.Save(headLink)
	if err != nil {
		return err
	}

	// mine the head link:
	results, err := app.minerApplication.Link(headLink.Hash().String())
	if err != nil {
		return err
	}

	// build the head mined link:
	head, err := app.minedLinkBuilder.Create().WithLink(headLink).WithMining(results).Now()
	if err != nil {
		return err
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
	chain, err := app.chainBuilder.Create().WithGenesis(gen).WithRoot(root).WithHead(head).WithTotal(totalAmount).Now()
	if err != nil {
		return err
	}

	// save the chain:
	return app.chainService.Insert(chain)
}

// Block mines a new block on the chain
func (app *application) Block(additional uint) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	// retrieve the genesis:
	gen, err := app.genesisRepository.Retrieve()
	if err != nil {
		return err
	}

	// retrieve the to-mine buckets from identity:
	toMinePrivBuckets := identity.Wallet().Miner().ToTransact().All()

	// create the to-mine buckets list:
	toMineBuckets := []buckets.Bucket{}
	for _, oneToMinePrivBucket := range toMinePrivBuckets {
		toMineBuckets = append(toMineBuckets, oneToMinePrivBucket.Bucket())
	}

	// create a new block:
	block, err := app.blockBuilder.Create().WithGenesis(gen).WithAdditional(additional).WithBuckets(toMineBuckets).Now()
	if err != nil {
		return err
	}

	// save the new block:
	err = app.blockService.Save(block)
	if err != nil {
		return err
	}

	// mine the block:
	results, err := app.minerApplication.Block(block.Hash().String())
	if err != nil {
		return err
	}

	// create a new mined block:
	minedBlock, err := app.minedBlockBuilder.Create().WithBlock(block).WithMining(results).Now()
	if err != nil {
		return err
	}

	// save the mined block to the identity for future lik mining:
	err = identity.Wallet().Miner().ToLink().Add(minedBlock)
	if err != nil {
		return err
	}

	for _, oneToMinePrivBucket := range toMinePrivBuckets {
		// delete the to-mine buckets from identity:
		err = identity.Wallet().Miner().ToTransact().Delete(oneToMinePrivBucket.Hash())
		if err != nil {
			return err
		}

		// add the private bucket in the queue list:
		err = identity.Wallet().Miner().Queue().Add(oneToMinePrivBucket)
		if err != nil {
			return err
		}
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

// Link mines a new link on the chain
func (app *application) Link(additional uint) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	// retrieve chain:
	chain, err := app.chainRepository.Retrieve()
	if err != nil {
		return err
	}

	// retrieve the to-mine blocks:
	gen := chain.Genesis()
	root := chain.Root()
	height := chain.Height()
	prev := chain.Head().Hash()
	toMineBlocks := identity.Wallet().Miner().ToLink().All()
	for _, oneToMineBlock := range toMineBlocks {
		// build a link:
		toMineLink, err := app.linkBuilder.Create().WithPrevious(prev).WithNext(oneToMineBlock).WithIndex(height + 1).Now()
		if err != nil {
			return err
		}

		// save the link:
		err = app.linkService.Save(toMineLink)
		if err != nil {
			return err
		}

		// mine the link:
		results, err := app.minerApplication.Link(toMineLink.Hash().String())
		if err != nil {
			return err
		}

		// build the mined link:
		minedLink, err := app.minedLinkBuilder.Create().WithLink(toMineLink).WithMining(results).Now()
		if err != nil {
			return err
		}

		// build an updated chain:
		blk := oneToMineBlock.Block()
		total := chain.Total() + uint(len(blk.Buckets())) + blk.Additional()
		updatedChain, err := app.chainBuilder.Create().WithGenesis(gen).WithRoot(root).WithHead(minedLink).WithTotal(total).Now()
		if err != nil {
			return err
		}

		// update the chain with the new mined link:
		err = app.chainService.Update(chain, updatedChain)
		if err != nil {
			return err
		}

		// remove the to-mine block from identity:
		err = identity.Wallet().Miner().ToLink().Delete(oneToMineBlock.Hash())
		if err != nil {
			return err
		}

		// save the buckets of the to-mine block as broadcasted:
		buckets := blk.Buckets()
		queuedPrivBuckets := identity.Wallet().Miner().Queue().All()
		for _, oneQueuedBucket := range queuedPrivBuckets {
			for _, oneBucket := range buckets {
				if !oneQueuedBucket.Bucket().Hash().Compare(oneBucket.Hash()) {
					continue
				}

				// delete the bucket from the queue:
				err = identity.Wallet().Miner().Queue().Delete(oneQueuedBucket.Hash())
				if err != nil {
					return err
				}

				// add the queued bucket as broadcasted:
				err = identity.Wallet().Miner().Broadcasted().Add(oneQueuedBucket)
				if err != nil {
					return err
				}

				break
			}
		}
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

func (app *application) mineEmptyBlock(gen genesis.Genesis, additional uint) (mined_block.Block, error) {
	// build the block:
	block, err := app.blockBuilder.Create().WithGenesis(gen).WithAdditional(additional).Now()
	if err != nil {
		return nil, err
	}

	// save the block:
	err = app.blockService.Save(block)
	if err != nil {
		return nil, err
	}

	// mine the block:
	results, err := app.minerApplication.Block(block.Hash().String())
	if err != nil {
		return nil, err
	}

	return app.minedBlockBuilder.Create().WithBlock(block).WithMining(results).Now()
}
