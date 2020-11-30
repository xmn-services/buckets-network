package identities

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	hashAdapter           hash.Adapter
	fileRepositoryBuilder file.EncryptedFileDiskRepositoryBuilder
	basePath              string
	extension             string
}

func createRepository(
	hashAdapter hash.Adapter,
	fileRepositoryBuilder file.EncryptedFileDiskRepositoryBuilder,
	basePath string,
	extension string,
) Repository {
	out := repository{
		hashAdapter:           hashAdapter,
		fileRepositoryBuilder: fileRepositoryBuilder,
		basePath:              basePath,
		extension:             extension,
	}

	return &out
}

// Retrieve retrieves an identity
func (app *repository) Retrieve(name string, seed string, password string) (Identity, error) {
	pass, err := makePassword(app.hashAdapter, seed, password)
	if err != nil {
		return nil, err
	}

	fileRepository, err := app.fileRepositoryBuilder.Create().WithBasePath(app.basePath).WithPassword(pass).Now()
	if err != nil {
		return nil, err
	}

	filePath := makeFileName(name, app.extension)
	js, err := fileRepository.Retrieve(filePath)
	if err != nil {
		return nil, err
	}

	ins := new(identity)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
