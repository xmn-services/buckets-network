package wallets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages"
)

// JSONWallet represents a JSON wallet instance
type JSONWallet struct {
	Miner    *miners.JSONMiner      `json:"miner"`
	Storage  *storages.JSONStorage  `json:"storage"`
	Accesses *accesses.JSONAccesses `json:"access"`
	Lists    *lists.JSONLists       `json:"lists"`
}

func createJSONWalletFromWallet(ins Wallet) *JSONWallet {
	minerAdapter := miners.NewAdapter()
	miner := minerAdapter.ToJSON(ins.Miner())

	storageAdapter := storages.NewAdapter()
	storage := storageAdapter.ToJSON(ins.Storage())

	accessesAdapter := accesses.NewAdapter()
	accesses := accessesAdapter.ToJSON(ins.Accesses())

	listsAdapter := lists.NewAdapter()
	lists := listsAdapter.ToJSON(ins.Lists())

	return createJSONWallet(miner, storage, accesses, lists)
}

func createJSONWallet(
	miner *miners.JSONMiner,
	storage *storages.JSONStorage,
	accesses *accesses.JSONAccesses,
	lists *lists.JSONLists,
) *JSONWallet {
	out := JSONWallet{
		Miner:    miner,
		Storage:  storage,
		Accesses: accesses,
		Lists:    lists,
	}

	return &out
}
