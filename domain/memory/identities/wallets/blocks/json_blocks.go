package blocks

import (
	"time"

	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
)

// JSONBlocks represents a JSON blocks instance
type JSONBlocks struct {
	Hash          string                    `json:"hash"`
	Blocks        []*mined_blocks.JSONBlock `json:"blocks"`
	CreatedOn     time.Time                 `json:"created_on"`
	LastUpdatedOn time.Time                 `json:"last_updated_on"`
}

func createJSONBlocksFromBlocks(blocks Blocks) *JSONBlocks {
	blocksAdapter := mined_blocks.NewAdapter()
	allBlocks := blocks.All()
	jsBlocks := []*mined_blocks.JSONBlock{}
	for _, oneBlock := range allBlocks {
		jsBlock := blocksAdapter.ToJSON(oneBlock)
		jsBlocks = append(jsBlocks, jsBlock)
	}

	hsh := blocks.Hash().String()
	createdOn := blocks.CreatedOn()
	lastUpdatedOn := blocks.LastUpdatedOn()
	return createJSONBlocks(hsh, jsBlocks, createdOn, lastUpdatedOn)
}

func createJSONBlocks(
	hash string,
	blocks []*mined_blocks.JSONBlock,
	createdOn time.Time,
	lastUpdatedOn time.Time,
) *JSONBlocks {
	out := JSONBlocks{
		Hash:          hash,
		Blocks:        blocks,
		CreatedOn:     createdOn,
		LastUpdatedOn: lastUpdatedOn,
	}

	return &out
}
