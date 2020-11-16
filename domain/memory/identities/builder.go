package identities

import (
	"errors"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets"
)

type builder struct {
	walletFactory wallets.Factory
	seed          string
	name          string
	root          string
	wallet        wallets.Wallet
}

func createBuilder(
	walletFactory wallets.Factory,
) Builder {
	out := builder{
		walletFactory: walletFactory,
		seed:          "",
		name:          "",
		root:          "",
		wallet:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
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

	if app.wallet == nil {
		wallet, err := app.walletFactory.Create()
		if err != nil {
			return nil, err
		}

		app.wallet = wallet
	}

	return createIdentity(app.seed, app.name, app.root, app.wallet), nil
}
