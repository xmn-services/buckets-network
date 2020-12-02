package files

import (
	"io/ioutil"
	"math"
	"path/filepath"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	transfer_file "github.com/xmn-services/buckets-network/domain/transfers/buckets/files"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	hashAdapter     hash.Adapter
	chunkRepository chunks.Repository
	trRepository    transfer_file.Repository
	builder         Builder
	chunkBuilder    chunks.Builder
	chkSizeInBytes  uint
}

func createRepository(
	hashAdapter hash.Adapter,
	chunkRepository chunks.Repository,
	trRepository transfer_file.Repository,
	builder Builder,
	chkSizeInBytes uint,
) Repository {
	out := repository{
		hashAdapter:     hashAdapter,
		chunkRepository: chunkRepository,
		trRepository:    trRepository,
		builder:         builder,
		chkSizeInBytes:  chkSizeInBytes,
	}

	return &out
}

// Retrieve retrieves a file instance by hash
func (app *repository) Retrieve(hsh hash.Hash) (File, error) {
	trFile, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	amount := trFile.Amount()
	chunkHashes := []hash.Hash{}
	leaves := trFile.Chunks().Parent().BlockLeaves().Leaves()
	for i := 0; i < int(amount); i++ {
		chunkHashes = append(chunkHashes, leaves[i].Head())
	}

	chunks, err := app.chunkRepository.RetrieveAll(chunkHashes)
	if err != nil {
		return nil, err
	}

	sizeInBytes := trFile.RelativePath()
	createdOn := trFile.CreatedOn()
	return app.builder.Create().WithRelativePath(sizeInBytes).WithChunks(chunks).CreatedOn(createdOn).Now()
}

// RetrieveAll retrieves all file instances by hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]File, error) {
	out := []File{}
	for _, oneHash := range hashes {
		chk, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, chk)
	}

	return out, nil
}

// RetrieveAllWithChunksContentFromPath retrieves files with chunk's contents from path
func (app *repository) RetrieveAllWithChunksContentFromPath(path string, pubKey public.Key) ([]File, [][][]byte, error) {
	return app.dirToFiles(path, ".", pubKey)
}

func (app *repository) dirToFiles(rootPath string, relativePath string, pubKey public.Key) ([]File, [][][]byte, error) {
	path := filepath.Join(rootPath, relativePath)
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	out := []File{}
	chks := [][][]byte{}
	for _, oneFile := range dirFiles {
		name := oneFile.Name()
		filePath := filepath.Join(relativePath, name)
		if oneFile.IsDir() {
			subFiles, subChksData, err := app.dirToFiles(rootPath, filePath, pubKey)
			if err != nil {
				return nil, nil, err
			}

			out = append(out, subFiles...)
			chks = append(chks, subChksData...)
			continue
		}

		file, chksData, err := app.dirFileToFile(rootPath, filePath, pubKey)
		if err != nil {
			return nil, nil, err
		}

		out = append(out, file)
		chks = append(chks, chksData)
	}

	return out, chks, nil
}

func (app *repository) dirFileToFile(rootPath string, relativePath string, pubKey public.Key) (File, [][]byte, error) {
	path := filepath.Join(rootPath, relativePath)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	index := 0
	chkData := [][]byte{}
	chunks := []chunks.Chunk{}
	loops := int(math.Ceil(float64(len(data)) / float64(app.chkSizeInBytes)))
	for i := 0; i < loops; i++ {
		beginsOn := i * int(app.chkSizeInBytes)
		createdOn := time.Now().UTC()
		if (i + 1) == loops {
			dataChk, dataHash, sizeInBytes, err := app.transformData(data[beginsOn:], pubKey)
			if err != nil {
				return nil, nil, err
			}

			chk, err := app.chunkBuilder.Create().WithSizeInBytes(sizeInBytes).WithData(*dataHash).CreatedOn(createdOn).Now()
			if err != nil {
				return nil, nil, err
			}

			chunks = append(chunks, chk)
			chkData = append(chkData, dataChk)
			continue
		}

		stopsOn := (i + 1) * int(app.chkSizeInBytes)
		dataChk := data[beginsOn:stopsOn]
		sizeInBytes := len(dataChk)
		dataHash, err := app.hashAdapter.FromBytes(dataChk)
		if err != nil {
			return nil, nil, err
		}

		chk, err := app.chunkBuilder.Create().WithSizeInBytes(uint(sizeInBytes)).WithData(*dataHash).CreatedOn(createdOn).Now()
		if err != nil {
			return nil, nil, err
		}

		chunks = append(chunks, chk)
		chkData = append(chkData, dataChk)
		index++
	}

	createdOn := time.Now().UTC()
	file, err := app.builder.Create().WithRelativePath(relativePath).WithChunks(chunks).CreatedOn(createdOn).Now()
	if err != nil {
		return nil, nil, err
	}

	return file, chkData, nil
}

func (app *repository) transformData(input []byte, pubKey public.Key) ([]byte, *hash.Hash, uint, error) {
	if pubKey != nil {
		data, err := pubKey.Encrypt(input)
		if err != nil {
			return nil, nil, 0, err
		}

		input = data
	}

	hsh, err := app.hashAdapter.FromBytes(input)
	if err != nil {
		return nil, nil, 0, err
	}

	return input, hsh, uint(len(input)), nil
}
