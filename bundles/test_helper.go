package bundles

import (
	"os"

	"github.com/xmn-services/buckets-network/application/commands"
)

// CreateCommandApplicationForTests creates a new command application for tests
func CreateCommandApplicationForTests() commands.Application {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	peerFileNameWithExt := "peers.json"
	genesisFileNameWithExt := "genesis.json"
	chainFileName := "root"
	chainFileExt := "json"
	identityExt := "identity"
	chunkSizeInBytes := uint(1024 * 1024)
	encPKBitrate := 4096

	return NewCommandApplication(
		basePath,
		peerFileNameWithExt,
		genesisFileNameWithExt,
		chainFileName,
		chainFileExt,
		identityExt,
		chunkSizeInBytes,
		encPKBitrate,
	)
}
