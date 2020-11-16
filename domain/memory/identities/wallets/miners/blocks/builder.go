package blocks

import (
	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
)

type builder struct {
	lst []mined_blocks.Block
}

func createBuilder() Builder {
	out := builder{
		lst: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithBlocks add blocks to the builder
func (app *builder) WithBlocks(blocks []mined_blocks.Block) Builder {
	app.lst = blocks
	return app
}

// Now builds a new Blocks instance
func (app *builder) Now() (Blocks, error) {
	if app.lst == nil {
		app.lst = []mined_blocks.Block{}
	}

	mp := map[string]mined_blocks.Block{}
	for _, oneBlock := range app.lst {
		keyname := oneBlock.Hash().String()
		mp[keyname] = oneBlock
	}

	return createBlocks(app.lst, mp), nil
}
