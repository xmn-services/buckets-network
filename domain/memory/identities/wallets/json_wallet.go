package wallets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages"
)

// JSONWallet represents a JSON wallet instance
type JSONWallet struct {
	Miner   *miners.JSONMiner     `json:"miner"`
	Storage *storages.JSONStorage `json:"storage"`
}

func createJSONWalletFromWallet(ins Wallet) *JSONWallet {
	minerAdapter := miners.NewAdapter()
	miner := minerAdapter.ToJSON(ins.Miner())

	storageAdapter := storages.NewAdapter()
	storage := storageAdapter.ToJSON(ins.Storage())

	return createJSONWallet(miner, storage)
}

func createJSONWallet(
	miner *miners.JSONMiner,
	storage *storages.JSONStorage,
) *JSONWallet {
	out := JSONWallet{
		Miner:   miner,
		Storage: storage,
	}

	return &out
}
