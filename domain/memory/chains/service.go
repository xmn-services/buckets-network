package chains

import (
	"errors"
	"fmt"
	"strings"

	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	transfer_chains "github.com/xmn-services/buckets-network/domain/transfers/chains"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type service struct {
	hashAdapter    hash.Adapter
	adapter        Adapter
	repository     Repository
	genesisService genesis.Service
	blockService   mined_block.Service
	linkRepository mined_link.Repository
	linkService    mined_link.Service
	trService      transfer_chains.Service
}

func createService(
	hashAdapter hash.Adapter,
	adapter Adapter,
	repository Repository,
	genesisService genesis.Service,
	blockService mined_block.Service,
	linkRepository mined_link.Repository,
	linkService mined_link.Service,
	trService transfer_chains.Service,
) Service {
	out := service{
		hashAdapter:    hashAdapter,
		adapter:        adapter,
		repository:     repository,
		genesisService: genesisService,
		blockService:   blockService,
		linkRepository: linkRepository,
		linkService:    linkService,
		trService:      trService,
	}

	return &out
}

// Update updates a chain
func (app *service) Update(original Chain, updated Chain) error {
	// make sure the the genesis is the same in both chains:
	updatedGenHash := updated.Genesis().Hash()
	originalGenHash := original.Genesis().Hash()
	if originalGenHash.Compare(updatedGenHash) {
		str := fmt.Sprintf("the chain cannot be updated at height (%d) because its Genesis instance is invalid (updated: %s, stored: %s)", original.Head().Link().Index(), updatedGenHash.String(), originalGenHash.String())
		return errors.New(str)
	}

	// make sure the root is the same in both chains:
	updatedRootHash := updated.Root().Hash()
	originalRootHash := original.Root().Hash()
	if originalRootHash.Compare(updatedRootHash) {
		str := fmt.Sprintf("the chain cannot be updated at height (%d) because its Root mined Block instance is invalid (updated: %s, stored: %s)", original.Head().Link().Index(), updatedRootHash.String(), originalRootHash.String())
		return errors.New(str)
	}

	return app.save(updated)
}

// Insert inserts a chain
func (app *service) Insert(chain Chain) error {
	_, err := app.repository.Retrieve()
	if err == nil {
		return nil
	}

	// retrieve data:
	gen := chain.Genesis()
	root := chain.Root()

	// save genesis:
	err = app.genesisService.Save(gen)
	if err != nil {
		return err
	}

	// save root:
	err = app.blockService.Save(root)
	if err != nil {
		return err
	}

	return app.save(chain)
}

func (app *service) save(chain Chain) error {
	// verify the chain:
	err := app.verifyChain(chain)
	if err != nil {
		return err
	}

	// retrieve data:
	head := chain.Head()

	// save head:
	err = app.linkService.Save(head)
	if err != nil {
		return err
	}

	// save the transfer chain:
	trChain, err := app.adapter.ToTransfer(chain)
	if err != nil {
		return err
	}

	return app.trService.Save(trChain)
}

func (app *service) verifyChain(chain Chain) error {
	gen := chain.Genesis()

	// verify the root block:
	root := chain.Root()
	err := app.verifyBlock(gen, root)
	if err != nil {
		return err
	}

	// verify the head link:
	head := chain.Head()
	expectedIndex := chain.Height()
	bucketsAmount, err := app.verifyLink(gen, expectedIndex, head)
	if err != nil {
		return err
	}

	totalBucketsAmount := bucketsAmount + uint(len(root.Block().Buckets()))
	if totalBucketsAmount != chain.Total() {
		str := fmt.Sprintf("the chain (hash: %s) was expecting %d buckets, %d found", chain.Hash().String(), totalBucketsAmount, chain.Total())
		return errors.New(str)
	}

	return nil
}

func (app *service) verifyLink(chainGen genesis.Genesis, expectedIndex uint, minedLink mined_link.Link) (uint, error) {
	link := minedLink.Link()
	block := link.Next().Block()
	gen := link.Next().Block().Genesis()
	if chainGen.Hash().Compare(gen.Hash()) {
		str := fmt.Sprintf("the chain has a different genesis (%s) than the mined link (hash: %s) genesis (%s) at index: %d", chainGen.Hash().String(), minedLink.Hash().String(), gen.Hash().String(), minedLink.Link().Index())
		return 0, errors.New(str)
	}

	if expectedIndex != link.Index() {
		str := fmt.Sprintf("the mined link (hash: %s) index (%d) was expected to be %d", minedLink.Hash().String(), link.Index(), expectedIndex)
		return 0, errors.New(str)
	}

	miningValue := gen.MiningValue()
	diff := gen.Difficulty().Link()
	linkMining := minedLink.Mining()
	linkhash := minedLink.Hash()

	// verify the mining:
	err := app.verifyMining(miningValue, diff, linkMining, linkhash)
	if err != nil {
		return 0, err
	}

	prevLinkHash := link.PreviousLink()
	prevLink, err := app.linkRepository.Retrieve(prevLinkHash)
	if err != nil {
		return 0, err
	}

	prevAmount, err := app.verifyLink(chainGen, expectedIndex-1, prevLink)
	if err != nil {
		return 0, err
	}

	additional := block.Additional()
	amount := uint(len(block.Buckets()))
	return prevAmount + additional + amount, nil
}

func (app *service) fetchAmountBuckets(minedLink mined_link.Link) (uint, error) {
	link := minedLink.Link()
	block := link.Next().Block()
	additional := block.Additional()
	amount := uint(len(block.Buckets()))

	prevMinedLinkHash := link.PreviousLink()
	prevMinedLink, err := app.linkRepository.Retrieve(prevMinedLinkHash)
	if err != nil {
		return 0, err
	}

	prevMinedAmount, err := app.fetchAmountBuckets(prevMinedLink)
	if err != nil {
		return 0, err
	}

	return prevMinedAmount + additional + amount, nil

}

func (app *service) verifyBlock(chainGen genesis.Genesis, minedBlock mined_block.Block) error {
	block := minedBlock.Block()
	gen := block.Genesis()
	if chainGen.Hash().Compare(gen.Hash()) {
		str := fmt.Sprintf("the chain has a different genesis (%s) than the block (hash: %s) genesis (%s)", chainGen.Hash().String(), minedBlock.Hash().String(), gen.Hash().String())
		return errors.New(str)
	}

	blockMining := minedBlock.Mining()
	blockHash := block.Hash()
	amountBuckets := block.Additional() + uint(len(block.Buckets()))
	miningValue := gen.MiningValue()

	// calculate the difficulty:
	difficulty := gen.Difficulty()
	blockDifficulty := difficulty.Block()
	blockBaseDifficulty := blockDifficulty.Base()
	blockIncreaseDifficultyPerBucket := blockDifficulty.IncreasePerBucket()
	diff := BlockDifficulty(blockBaseDifficulty, blockIncreaseDifficultyPerBucket, amountBuckets)

	// verify the mining:
	err := app.verifyMining(miningValue, diff, blockMining, blockHash)
	if err != nil {
		return err
	}

	return nil
}

func (app *service) verifyMining(miningValue uint8, difficulty uint, mining string, hash hash.Hash) error {
	miningHash, err := app.hashAdapter.FromBytes([]byte(mining))
	if err != nil {
		return err
	}

	prefix := ""
	for i := 0; i < int(difficulty); i++ {
		prefix = fmt.Sprintf("%s%d", prefix, miningValue)
	}

	miningHashStr := miningHash.String()
	if strings.HasPrefix(miningHashStr, prefix) {
		return nil
	}

	str := fmt.Sprintf("the mining (%s), when hashed (%s) was expected to contain %d times the character '%d' as prefix in order to be valid", mining, miningHashStr, difficulty, miningValue)
	return errors.New(str)
}
