package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files"
)

// JSONStorage represents a JSON storage instance
type JSONStorage struct {
	ToDownload *files.JSONFiles `json:"to_download"`
	Stored     *files.JSONFiles `json:"stored"`
}

func createJSONStorageFromStorage(ins Storage) *JSONStorage {
	fileAdapter := files.NewAdapter()
	toDownload := fileAdapter.ToJSON(ins.ToDownload())
	stored := fileAdapter.ToJSON(ins.Stored())
	return createJSONStorage(toDownload, stored)
}

func createJSONStorage(
	toDownload *files.JSONFiles,
	stored *files.JSONFiles,
) *JSONStorage {
	out := JSONStorage{
		ToDownload: toDownload,
		Stored:     stored,
	}

	return &out
}
