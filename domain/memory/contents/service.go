package contents

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	transfer_content "github.com/xmn-services/buckets-network/domain/transfers/contents"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type service struct {
	hashAdapter      hash.Adapter
	repository       Repository
	trContentService transfer_content.Service
}

func createService(
	hashAdapter hash.Adapter,
	repository Repository,
	trContentService transfer_content.Service,
) Service {
	out := service{
		hashAdapter:      hashAdapter,
		repository:       repository,
		trContentService: trContentService,
	}

	return &out
}

// Extract extract the bucket's data to the given absolute path
func (app *service) Extract(bucket buckets.Bucket, decryptPrivKey encryption.PrivateKey, absolutePath string) error {
	files := bucket.Files()
	for _, oneFile := range files {
		fileHash := oneFile.Hash()
		relativePath := oneFile.RelativePath()
		fileContents := []byte{}
		chks := oneFile.Chunks()
		for _, oneChunk := range chks {
			chkHash := oneChunk.Data()
			content, err := app.repository.Retrieve(bucket, fileHash, chkHash)
			if err != nil {
				return err
			}

			fileContents = append(fileContents, content...)
		}

		// decrypt the content:
		decryptedFileContent, err := decryptPrivKey.Decrypt(fileContents)
		if err != nil {
			return err
		}

		// create directory if it doesn't exists:
		path := filepath.Join(absolutePath, relativePath)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			dir := filepath.Dir(path)
			os.MkdirAll(dir, 0777)
		}

		// save decrypted file on disk:
		err = ioutil.WriteFile(path, decryptedFileContent, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}

// Save saves a chunk in bucket
func (app *service) Save(bucket buckets.Bucket, data []byte) error {
	hsh, err := app.hashAdapter.FromBytes(data)
	if err != nil {
		return err
	}

	file, _, err := bucket.FileChunkByHash(*hsh)
	if err != nil {
		return err
	}

	return app.trContentService.Save(bucket.Hash(), file.Hash(), data)
}

// Delete deletes a chunk from bucket
func (app *service) Delete(bucket buckets.Bucket, chunkHash hash.Hash) error {
	file, chunk, err := bucket.FileChunkByHash(chunkHash)
	if err != nil {
		return err
	}

	return app.trContentService.Delete(bucket.Hash(), file.Hash(), chunk.Hash())
}

// DeleteAll deletes all chunks from bucket
func (app *service) DeleteAll(bucket buckets.Bucket) error {
	return app.trContentService.DeleteAll(bucket.Hash())
}
