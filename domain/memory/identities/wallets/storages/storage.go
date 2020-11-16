package storages

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files"
)

type storage struct {
	toDownload files.Files
	stored     files.Files
}

func createStorageFromJSON(ins *JSONStorage) (Storage, error) {
	fileAdapter := files.NewAdapter()
	toDownload, err := fileAdapter.ToFiles(ins.ToDownload)
	if err != nil {
		return nil, err
	}

	stored, err := fileAdapter.ToFiles(ins.Stored)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithToDownload(toDownload).
		WithStored(stored).
		Now()
}

func createStorage(
	toDownload files.Files,
	stored files.Files,
) Storage {
	out := storage{
		toDownload: toDownload,
		stored:     stored,
	}

	return &out
}

// ToDownload returns the toDownload files
func (obj *storage) ToDownload() files.Files {
	return obj.toDownload
}

// Stored returns the stored files
func (obj *storage) Stored() files.Files {
	return obj.stored
}

// MarshalJSON converts the instance to JSON
func (obj *storage) MarshalJSON() ([]byte, error) {
	ins := createJSONStorageFromStorage(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *storage) UnmarshalJSON(data []byte) error {
	ins := new(JSONStorage)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createStorageFromJSON(ins)
	if err != nil {
		return err
	}

	insBucket := pr.(*storage)
	obj.toDownload = insBucket.toDownload
	obj.stored = insBucket.stored
	return nil
}
