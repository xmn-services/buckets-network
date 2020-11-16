package identities

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets"
)

type identity struct {
	seed   string
	name   string
	root   string
	wallet wallets.Wallet
}

func createIdentityFromJSON(ins *JSONIdentity) (Identity, error) {
	walletAdapter := wallets.NewAdapter()
	wallet, err := walletAdapter.ToWallet(ins.Wallet)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithWallet(wallet).
		WithSeed(ins.Seed).
		WithName(ins.Name).
		WithRoot(ins.Root).
		Now()
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

// MarshalJSON converts the instance to JSON
func (obj *identity) MarshalJSON() ([]byte, error) {
	ins := createJSONIdentityFromIdentity(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *identity) UnmarshalJSON(data []byte) error {
	ins := new(JSONIdentity)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createIdentityFromJSON(ins)
	if err != nil {
		return err
	}

	insIdentity := pr.(*identity)
	obj.seed = insIdentity.seed
	obj.name = insIdentity.name
	obj.root = insIdentity.root
	obj.wallet = insIdentity.wallet
	return nil
}
