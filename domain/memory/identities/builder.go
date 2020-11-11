package identities

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	mutableBuilder entities.MutableBuilder
	walletFactory  wallets.Factory
	seed           string
	name           string
	root           string
	wallet         wallets.Wallet
	createdOn      *time.Time
	lastUpdatedOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	mutableBuilder entities.MutableBuilder,
	walletFactory wallets.Factory,
) Builder {
	out := builder{
		hashAdapter:    hashAdapter,
		mutableBuilder: mutableBuilder,
		walletFactory:  walletFactory,
		seed:           "",
		name:           "",
		root:           "",
		wallet:         nil,
		createdOn:      nil,
		lastUpdatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.mutableBuilder,
		app.walletFactory,
	)
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root string) Builder {
	app.root = root
	return app
}

// WithWallet adds a wallet to the builder
func (app *builder) WithWallet(wallet wallets.Wallet) Builder {
	app.wallet = wallet
	return app
}

// CreatedOn add a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// LastUpdatedOn add a lastUpdatedOn time to the builder
func (app *builder) LastUpdatedOn(lastUpdatedOn time.Time) Builder {
	app.lastUpdatedOn = &lastUpdatedOn
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Identity instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	if app.root == "" {
		return nil, errors.New("the root is mandatory in order to build an Identity instance")
	}

	data := [][]byte{
		[]byte(app.seed),
		[]byte(app.name),
		[]byte(app.root),
	}

	if app.wallet != nil {
		data = append(data, app.wallet.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	if app.wallet == nil {
		app.wallet = app.walletFactory.Create()
	}

	return createIdentity(mutable, app.seed, app.name, app.root, app.wallet), nil
}
