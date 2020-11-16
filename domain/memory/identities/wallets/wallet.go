package wallets

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages"
)

type wallet struct {
	miner   miners.Miner
	storage storages.Storage
}

func createWalletFromJSON(ins *JSONWallet) (Wallet, error) {
	minerAdapter := miners.NewAdapter()
	miner, err := minerAdapter.ToMiner(ins.Miner)
	if err != nil {
		return nil, err
	}

	storageAdapter := storages.NewAdapter()
	storage, err := storageAdapter.ToStorage(ins.Storage)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithMiner(miner).
		WithStorage(storage).
		Now()
}

func createWallet(
	miner miners.Miner,
	storage storages.Storage,
) Wallet {
	out := wallet{
		miner:   miner,
		storage: storage,
	}

	return &out
}

// Miner returns the miner
func (obj *wallet) Miner() miners.Miner {
	return obj.miner
}

// Storage returns the storage
func (obj *wallet) Storage() storages.Storage {
	return obj.storage
}

// MarshalJSON converts the instance to JSON
func (obj *wallet) MarshalJSON() ([]byte, error) {
	ins := createJSONWalletFromWallet(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *wallet) UnmarshalJSON(data []byte) error {
	ins := new(JSONWallet)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createWalletFromJSON(ins)
	if err != nil {
		return err
	}

	insWallet := pr.(*wallet)
	obj.miner = insWallet.miner
	obj.storage = insWallet.storage
	return nil
}
