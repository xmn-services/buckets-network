package identities

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type identity struct {
	mutable entities.Mutable
	seed    string
	name    string
	root    string
	wallet  wallets.Wallet
}

func createIdentity(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	wallet wallets.Wallet,
) Identity {
	return createIdentityInternally(mutable, seed, name, root, wallet)
}

func createIdentityInternally(
	mutable entities.Mutable,
	seed string,
	name string,
	root string,
	wallet wallets.Wallet,
) Identity {
	out := identity{
		mutable: mutable,
		seed:    seed,
		name:    name,
		root:    root,
		wallet:  wallet,
	}

	return &out
}

// Hash returns the hash
func (obj *identity) Hash() hash.Hash {
	return obj.mutable.Hash()
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

// LastUpdatedOn returns the lastUpdatedOn time
func (obj *identity) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}

// CreatedOn returns the creation time
func (obj *identity) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}
