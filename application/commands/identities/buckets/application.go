package buckets

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	identity_buckets "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter           hash.Adapter
	pkFactory             encryption.Factory
	chunkBuilder          chunks.Builder
	fileBuilder           files.Builder
	bucketBuilder         buckets.Builder
	bucketRepository      buckets.Repository
	bucketService         buckets.Service
	identityBucketBuilder identity_buckets.Builder
	identityRepository    identities.Repository
	identityService       identities.Service
	contentService        contents.Service
	name                  string
	password              string
	seed                  string
	chunkSizeInBytes      uint
}

func createApplication(
	hashAdapter hash.Adapter,
	pkFactory encryption.Factory,
	chunkBuilder chunks.Builder,
	fileBuilder files.Builder,
	bucketBuilder buckets.Builder,
	bucketRepository buckets.Repository,
	bucketService buckets.Service,
	identityBucketBuilder identity_buckets.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
	contentService contents.Service,
	name string,
	password string,
	seed string,
	chunkSizeInBytes uint,
) Application {
	out := application{
		hashAdapter:           hashAdapter,
		pkFactory:             pkFactory,
		chunkBuilder:          chunkBuilder,
		fileBuilder:           fileBuilder,
		bucketBuilder:         bucketBuilder,
		bucketRepository:      bucketRepository,
		bucketService:         bucketService,
		identityBucketBuilder: identityBucketBuilder,
		identityRepository:    identityRepository,
		identityService:       identityService,
		contentService:        contentService,
		name:                  name,
		password:              password,
		seed:                  seed,
		chunkSizeInBytes:      chunkSizeInBytes,
	}

	return &out
}

// Add adds the bucket path
func (app *application) Add(relativePath string) error {
	if relativePath == "" {
		return errors.New("the path cannot be empty")
	}

	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return err
	}

	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	pk, err := app.pkFactory.Create()
	if err != nil {
		return err
	}

	pubKey := pk.Public()
	files, filesChksData, err := app.dirToFiles(absolutePath, ".", pubKey)
	if err != nil {
		return err
	}

	bucketCreatedOn := time.Now().UTC()
	bucket, err := app.bucketBuilder.Create().WithFiles(files).CreatedOn(bucketCreatedOn).Now()
	if err != nil {
		return err
	}

	// save the bucket:
	err = app.bucketService.Save(bucket)
	if err != nil {
		return err
	}

	// for each file, fetch the chunks and save its data:
	for _, oneFileChkData := range filesChksData {
		for _, oneChkData := range oneFileChkData {
			err := app.contentService.Save(bucket, oneChkData)
			if err != nil {
				return err
			}
		}
	}

	createdOn := time.Now().UTC()
	identityBucket, err := app.identityBucketBuilder.Create().
		WithBucket(bucket).
		WithAbsolutePath(absolutePath).
		WithPrivateKey(pk).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return err
	}

	err = identity.Wallet().Miner().ToTransact().Add(identityBucket)
	if err != nil {
		return err
	}

	return app.identityService.Update(identity, app.password, app.password)
}

// Delete deletes a bucket from the given path
func (app *application) Delete(hashStr string) error {
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	hash, err := app.hashAdapter.FromString(hashStr)
	if err != nil {
		return err
	}

	// retrieve the bucket:
	bucket, err := app.bucketRepository.Retrieve(*hash)
	if err != nil {
		return err
	}

	// delete the bucket:
	err = app.bucketService.Delete(bucket)
	if err != nil {
		return err
	}

	// delete the bucket's content:
	err = app.contentService.DeleteAll(bucket)
	if err != nil {
		return err
	}

	err = identity.Wallet().Miner().ToTransact().Delete(*hash)
	if err != nil {
		return err
	}

	return app.identityService.Update(identity, app.password, app.password)
}

// Retrieve retrieves a bucket by hash
func (app *application) Retrieve(hashStr string) (buckets.Bucket, error) {
	hash, err := app.hashAdapter.FromString(hashStr)
	if err != nil {
		return nil, err
	}

	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return nil, err
	}

	if !identity.Wallet().Storage().Stored().Exists(*hash) {
		str := fmt.Sprintf("the bucket (hash: %s) does not exists", hash.String())
		return nil, errors.New(str)
	}

	return app.bucketRepository.Retrieve(*hash)
}

// RetrieveAll retrieves all the buckets
func (app *application) RetrieveAll() ([]buckets.Bucket, error) {
	fetchBuckets := func(identityBuckets []identity_buckets.Bucket) []buckets.Bucket {
		out := []buckets.Bucket{}
		for _, oneIdentityBucket := range identityBuckets {
			out = append(out, oneIdentityBucket.Bucket())
		}

		return out
	}

	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return nil, err
	}

	buckets := []buckets.Bucket{}
	identityMiner := identity.Wallet().Miner()
	toTransactBuckets := identityMiner.ToTransact().All()
	broadcastedBuckets := identityMiner.Broadcasted().All()
	queuedBuckets := identityMiner.Queue().All()

	buckets = append(buckets, fetchBuckets(toTransactBuckets)...)
	buckets = append(buckets, fetchBuckets(broadcastedBuckets)...)
	buckets = append(buckets, fetchBuckets(queuedBuckets)...)

	toLinkBlocks := identityMiner.ToLink().All()
	for _, oneToLinkBlock := range toLinkBlocks {
		block := oneToLinkBlock.Block()
		if block.HasBuckets() {
			toLinkBuckets := block.Buckets()
			buckets = append(buckets, toLinkBuckets...)
		}
	}

	bucketHashes := []hash.Hash{}
	contents := identity.Wallet().Storage().Stored().All()
	for _, oneContent := range contents {
		bucketHashes = append(bucketHashes, oneContent.Bucket())
	}

	storedBuckets, err := app.bucketRepository.RetrieveAll(bucketHashes)
	if err != nil {
		return nil, err
	}

	buckets = append(buckets, storedBuckets...)
	return buckets, nil
}

func (app *application) dirToFiles(rootPath string, relativePath string, pubKey public.Key) ([]files.File, [][][]byte, error) {
	path := filepath.Join(rootPath, relativePath)
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	out := []files.File{}
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

func (app *application) transformData(input []byte, pubKey public.Key) ([]byte, *hash.Hash, uint, error) {
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

func (app *application) dirFileToFile(rootPath string, relativePath string, pubKey public.Key) (files.File, [][]byte, error) {
	path := filepath.Join(rootPath, relativePath)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	index := 0
	chkData := [][]byte{}
	chunks := []chunks.Chunk{}
	loops := int(math.Ceil(float64(len(data)) / float64(app.chunkSizeInBytes)))
	for i := 0; i < loops; i++ {
		beginsOn := i * int(app.chunkSizeInBytes)
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

		stopsOn := (i + 1) * int(app.chunkSizeInBytes)
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
	file, err := app.fileBuilder.Create().WithRelativePath(relativePath).WithChunks(chunks).CreatedOn(createdOn).Now()
	if err != nil {
		return nil, nil, err
	}

	return file, chkData, nil
}
