package identities

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets"
)

// JSONIdentity represents a JSON identity instance
type JSONIdentity struct {
	Seed   string              `json:"seed"`
	Name   string              `json:"name"`
	Root   string              `json:"root"`
	Wallet *wallets.JSONWallet `json:"wallet"`
}

func createJSONIdentityFromIdentity(ins Identity) *JSONIdentity {
	walletAdapter := wallets.NewAdapter()
	wallet := walletAdapter.ToJSON(ins.Wallet())

	seed := ins.Seed()
	name := ins.Name()
	root := ins.Root()

	return createJSONIdentity(seed, name, root, wallet)
}

func createJSONIdentity(
	seed string,
	name string,
	root string,
	wallet *wallets.JSONWallet,
) *JSONIdentity {
	out := JSONIdentity{
		Seed:   seed,
		Name:   name,
		Root:   root,
		Wallet: wallet,
	}

	return &out
}
