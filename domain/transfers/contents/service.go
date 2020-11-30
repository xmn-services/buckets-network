package contents

import (
	"path/filepath"

	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type service struct {
	hashAdapter hash.Adapter
	fileService file.Service
}

func createService(
	hashAdapter hash.Adapter,
	fileService file.Service,
) Service {
	out := service{
		hashAdapter: hashAdapter,
		fileService: fileService,
	}

	return &out
}

// Save saves content in bucket
func (app *service) Save(bucketHash hash.Hash, fileHash hash.Hash, data []byte) error {
	chunkHash, err := app.hashAdapter.FromBytes(data)
	if err != nil {
		return err
	}

	path := filepath.Join(bucketHash.String(), fileHash.String(), chunkHash.String())
	return app.fileService.Save(path, data)
}

// Delete deletes a file's chunk from bucket
func (app *service) Delete(bucketHash hash.Hash, fileHash hash.Hash, chunkHash hash.Hash) error {
	path := filepath.Join(bucketHash.String(), fileHash.String(), chunkHash.String())
	return app.fileService.Delete(path)
}

// DeleteFile deletes a file from bucket
func (app *service) DeleteFile(bucketHash hash.Hash, fileHash hash.Hash) error {
	path := filepath.Join(bucketHash.String(), fileHash.String())
	return app.fileService.Delete(path)
}

// DeleteAll deletes all files from bucket
func (app *service) DeleteAll(bucketHash hash.Hash) error {
	return app.fileService.Delete(bucketHash.String())
}
