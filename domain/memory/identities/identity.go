package identities

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets"
)

type identity struct {
	seed   string
	name   string
	root   string
	wallet wallets.Wallet
}

func createIdentity(
	seed string,
	name string,
	root string,
	wallet wallets.Wallet,
) Identity {
	out := identity{
		seed:   seed,
		name:   name,
		root:   root,
		wallet: wallet,
	}

	return &out
}

// Seed returns the seed
func (obj *identity) Seed() string {
	return obj.seed
}

// SetSeed sets a seed
func (obj *identity) SetSeed(seed string) {
	obj.seed = seed
}

// Name returns the name
func (obj *identity) Name() string {
	return obj.name
}

// SetName sets a name
func (obj *identity) SetName(name string) {
	obj.name = name
}

// Root returns the root
func (obj *identity) Root() string {
	return obj.root
}

// SetRoot sets a root
func (obj *identity) SetRoot(root string) {
	obj.root = root
}

// Wallet returns the wallet
func (obj *identity) Wallet() wallets.Wallet {
	return obj.wallet
}
