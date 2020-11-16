package blocks

import (
	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
)

// JSONBlocks represents a JSON blocks instance
type JSONBlocks struct {
	Blocks []*mined_blocks.JSONBlock `json:"blocks"`
}

func createJSONBlocksFromBlocks(blocks Blocks) *JSONBlocks {
	blocksAdapter := mined_blocks.NewAdapter()
	allBlocks := blocks.All()
	jsBlocks := []*mined_blocks.JSONBlock{}
	for _, oneBlock := range allBlocks {
		jsBlock := blocksAdapter.ToJSON(oneBlock)
		jsBlocks = append(jsBlocks, jsBlock)
	}

	return createJSONBlocks(jsBlocks)
}

func createJSONBlocks(
	blocks []*mined_blocks.JSONBlock,
) *JSONBlocks {
	out := JSONBlocks{
		Blocks: blocks,
	}

	return &out
}
