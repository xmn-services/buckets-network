package file

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	transfer_data "github.com/xmn-services/buckets-network/domain/transfers/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	fileRepository   files.Repository
	trDataRepository transfer_data.Repository
	builder          Builder
}

func createRepository(
	fileRepository files.Repository,
	trDataRepository transfer_data.Repository,
	builder Builder,
) Repository {
	out := repository{
		fileRepository:   fileRepository,
		trDataRepository: trDataRepository,
		builder:          builder,
	}

	return &out
}

// Retrieve retrieves a file by hash
func (app *repository) Retrieve(fileHash hash.Hash) (File, error) {
	file, err := app.fileRepository.Retrieve(fileHash)
	if err != nil {
		return nil, err
	}

	contents := [][]byte{}
	chunks := file.Chunks()
	for _, oneChunk := range chunks {
		dataHash := oneChunk.Data()
		content, err := app.trDataRepository.Retrieve(dataHash)
		if err != nil {
			return nil, err
		}

		contents = append(contents, content)
	}

	builder := app.builder.Create().WithFile(file)
	if len(contents) > 0 {
		builder.WithContents(contents)
	}

	return builder.Now()
}
