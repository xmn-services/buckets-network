package files

import (
	transfer_file "github.com/xmn-services/buckets-network/domain/transfers/buckets/files"
	"github.com/xmn-services/buckets-network/libs/hashtree"
)

type adapter struct {
	hashTreeBuilder hashtree.Builder
	trBuilder       transfer_file.Builder
}

func createAdapter(
	hashTreeBuilder hashtree.Builder,
	trBuilder transfer_file.Builder,
) Adapter {
	out := adapter{
		hashTreeBuilder: hashTreeBuilder,
		trBuilder:       trBuilder,
	}

	return &out
}

// ToTransfer converts a file to a transfer file
func (app *adapter) ToTransfer(file File) (transfer_file.File, error) {
	hash := file.Hash()
	relativePath := file.RelativePath()
	chunks := file.Chunks()

	blocks := [][]byte{}
	for _, oneChunk := range chunks {
		blocks = append(blocks, oneChunk.Hash().Bytes())
	}

	ht, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
	if err != nil {
		return nil, err
	}

	amount := uint(len(chunks))
	createdOn := file.CreatedOn()
	return app.trBuilder.Create().
		WithHash(hash).
		WithRelativePath(relativePath).
		WithChunks(ht).
		WithAmount(amount).
		CreatedOn(createdOn).
		Now()
}

// ToJSON converts a file to JSON
func (app *adapter) ToJSON(file File) *JSONFile {
	return createJSONFileFromFile(file)
}

// ToFile converts JSON to file
func (app *adapter) ToFile(ins *JSONFile) (File, error) {
	return createFileFromJSON(ins)
}
