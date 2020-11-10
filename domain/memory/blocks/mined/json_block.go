package mined

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/blocks"
)

// JSONBlock represents a JSON block instance
type JSONBlock struct {
	Block     *blocks.JSONBlock `json:"block"`
	Mining    string            `json:"mining"`
	CreatedOn time.Time         `json:"created_on"`
}

func createJSONBlockFromBlock(block Block) *JSONBlock {
	blockAdapter := blocks.NewAdapter()
	jsBlock := blockAdapter.ToJSON(block.Block())

	mining := block.Mining()
	createdOn := block.CreatedOn()
	return createJSONBlock(jsBlock, mining, createdOn)
}

func createJSONBlock(
	block *blocks.JSONBlock,
	mining string,
	createdOn time.Time,
) *JSONBlock {
	out := JSONBlock{
		Block:     block,
		Mining:    mining,
		CreatedOn: createdOn,
	}

	return &out
}
