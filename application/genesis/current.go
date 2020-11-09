package genesis

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
)

type current struct {
	genesisBuilder     genesis.Builder
	genesisRepository  genesis.Repository
	genesisService     genesis.Service
	identityRepository identities.Repository
	identityBuilder    identities.Builder
	identityService    identities.Service
}

func createCurrent(
	genesisBuilder genesis.Builder,
	genesisRepository genesis.Repository,
	genesisService genesis.Service,
	identityRepository identities.Repository,
	identityBuilder identities.Builder,
	identityService identities.Service,
) Current {
	out := current{
		genesisBuilder:     genesisBuilder,
		genesisRepository:  genesisRepository,
		genesisService:     genesisService,
		identityRepository: identityRepository,
		identityBuilder:    identityBuilder,
		identityService:    identityService,
	}

	return &out
}

// Init initializes the genesis block
func (app *current) Init(
	name string,
	password string,
	seed string,
	walletName string,
	amountUnits uint64,
	blockDifficultyBase uint,
	blockDifficultyIncreasePerTrx float64,
	linkDifficulty uint,
) error {
	_, err := app.genesisRepository.Retrieve()
	if err == nil {
		return errors.New("the genesis block has already been created")
	}

	identity, err := app.identityRepository.Retrieve(name, seed, password)
	if err != nil {
		return err
	}

	createdOn := time.Now().UTC()
	gen, err := app.genesisBuilder.Create().
		WithBlockDifficultyBase(blockDifficultyBase).
		WithBlockDifficultyIncreasePerTrx(blockDifficultyIncreasePerTrx).
		WithLinkDifficulty(linkDifficulty).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return err
	}

	root := identity.Root()
	identityCreatedOn := identity.CreatedOn()
	lastUpdatedOn := time.Now().UTC()

	buckets := identity.Buckets().All()
	updatedIdentity, err := app.identityBuilder.Create().
		WithSeed(seed).
		WithName(name).
		WithRoot(root).
		WithBuckets(buckets).
		CreatedOn(identityCreatedOn).
		LastUpdatedOn(lastUpdatedOn).
		Now()

	if err != nil {
		return err
	}

	err = app.identityService.Update(identity.Hash(), updatedIdentity, password, password)
	if err != nil {
		return err
	}

	return app.genesisService.Save(gen)
}
