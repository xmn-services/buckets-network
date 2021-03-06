package blocks

import (
	transfer_block "github.com/xmn-services/buckets-network/domain/transfers/blocks"
	"github.com/xmn-services/buckets-network/libs/hashtree"
)

type adapter struct {
	hashTreeBuilder hashtree.Builder
	trBuilder       transfer_block.Builder
}

func createAdapter(
	hashTreeBuilder hashtree.Builder,
	trBuilder transfer_block.Builder,
) Adapter {
	out := adapter{
		hashTreeBuilder: hashTreeBuilder,
		trBuilder:       trBuilder,
	}

	return &out
}

// ToTransfer converts the block to a transfer block
func (app *adapter) ToTransfer(block Block) (transfer_block.Block, error) {
	hsh := block.Hash()
	additional := block.Additional()
	amount := uint(0)
	createdOn := block.CreatedOn()
	builder := app.trBuilder.Create().
		WithHash(hsh).
		WithAdditional(additional).
		WithAmount(amount).
		CreatedOn(createdOn)

	if block.HasBuckets() {
		buckets := block.Buckets()
		blocks := [][]byte{}
		for _, oneBucket := range buckets {
			blocks = append(blocks, oneBucket.Hash().Bytes())
		}

		bucketsHt, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
		if err != nil {
			return nil, err
		}

		amount := uint(len(buckets))
		builder.WithBuckets(bucketsHt).WithAmount(amount)
	}

	return builder.Now()
}

// ToJSON converts a block to a JSON block
func (app *adapter) ToJSON(block Block) *JSONBlock {
	return createJSONBlockFromBlock(block)
}

// ToBlock converts a JSON block to a block
func (app *adapter) ToBlock(ins *JSONBlock) (Block, error) {
	return createBlockFromJSON(ins)
}
