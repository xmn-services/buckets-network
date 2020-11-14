package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type fileDiskRepository struct {
	basePath string
}

func createFileDiskRepository(basePath string) Repository {
	out := fileDiskRepository{
		basePath: basePath,
	}

	return &out
}

// Exists returns true if the file exists, false otherwise
func (app *fileDiskRepository) Exists(relativePath string) bool {
	path := filepath.Join(app.basePath, relativePath)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Retrieve retrieves data from file using its name
func (app *fileDiskRepository) Retrieve(relativePath string) ([]byte, error) {
	if !app.Exists(relativePath) {
		str := fmt.Sprintf("the file (path: %s) does not exists", relativePath)
		return nil, errors.New(str)
	}

	path := filepath.Join(app.basePath, relativePath)
	encrypted, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

// RetrieveAll retrieves all files in a given directory
func (app *fileDiskRepository) RetrieveAll(relativePath string) ([]string, error) {
	path := filepath.Join(app.basePath, relativePath)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	out := []string{}
	for _, oneFile := range files {
		if oneFile.IsDir() {
			continue
		}

		out = append(out, oneFile.Name())
	}

	return out, nil
}
