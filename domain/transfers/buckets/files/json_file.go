package files

import (
	"time"

	"github.com/xmn-services/buckets-network/libs/hashtree"
)

type jsonFile struct {
	Hash         string                `json:"hash"`
	RelativePath string                `json:"relative_path"`
	Chunks       *hashtree.JSONCompact `json:"chunks"`
	Amount       uint                  `json:"amount"`
	CreatedOn    time.Time             `json:"created_on"`
}

func createJSONFileFromFile(ins File) *jsonFile {
	hash := ins.Hash().String()
	relativePath := ins.RelativePath()
	chunks := hashtree.NewAdapter().ToJSON(ins.Chunks().Compact())
	amount := ins.Amount()
	createdOn := ins.CreatedOn()
	return createJSONFile(hash, relativePath, chunks, amount, createdOn)
}

func createJSONFile(
	hash string,
	relativePath string,
	chunks *hashtree.JSONCompact,
	amount uint,
	createdOn time.Time,
) *jsonFile {
	out := jsonFile{
		Hash:         hash,
		RelativePath: relativePath,
		Chunks:       chunks,
		Amount:       amount,
		CreatedOn:    createdOn,
	}

	return &out
}
