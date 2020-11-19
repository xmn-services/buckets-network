package miners

import (
	"fmt"
	"testing"

	libs_file "github.com/xmn-services/buckets-network/libs/file"
)

func TestInit_Success(t *testing.T) {
	miningValue := uint8(0)
	baseDifficulty := uint(1)
	increasePerBucket := float64(0.05)
	linkDifficulty := uint(2)
	rootAdditionalBuckets := uint(40)
	headAdditionalBuckets := uint(20)

	basePath := "./test_files"
	fileRepository := libs_file.NewFileDiskRepository(basePath)
	fileService := libs_file.NewFileDiskService(basePath)
	genesisFileNameWithExt := "genesis.json"
	defer func() {
		fileService.Delete(genesisFileNameWithExt)
	}()

	app := NewApplication(fileRepository, fileService, genesisFileNameWithExt)
	chain, err := app.Init(miningValue, baseDifficulty, increasePerBucket, linkDifficulty, rootAdditionalBuckets, headAdditionalBuckets)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fmt.Printf("\n->%s\n", chain.Head().Mining())
}
