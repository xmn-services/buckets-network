package chains

import (
	"errors"

	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
)

type builder struct {
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

func createBuilder(
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
) Builder {
	out := builder{
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
		name:               "",
		password:           "",
		seed:               "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.minerApplication,
		app.identityRepository,
		app.identityService,
		app.genesisBuilder,
		app.genesisRepository,
		app.genesisService,
		app.blockBuilder,
		app.blockService,
		app.minedBlockBuilder,
		app.linkBuilder,
		app.linkService,
		app.minedLinkBuilder,
		app.chainBuilder,
		app.chainRepository,
		app.chainService,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) Builder {
	app.password = password
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build an Application instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.minerApplication,
		app.identityRepository,
		app.identityService,
		app.genesisBuilder,
		app.genesisRepository,
		app.genesisService,
		app.blockBuilder,
		app.blockService,
		app.minedBlockBuilder,
		app.linkBuilder,
		app.linkService,
		app.minedLinkBuilder,
		app.chainBuilder,
		app.chainRepository,
		app.chainService,
		app.name,
		app.password,
		app.seed,
	), nil
}
