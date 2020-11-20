package peers

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xmn-services/buckets-network/libs/file"
)

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	fileName := "peers.json"
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService, fileName)
	service := NewService(fileRepositoryService, fileName)
	return repository, service
}

// TestCompare compare two peers instances
func TestCompare(t *testing.T, first Peers, second Peers) {
	js, err := json.Marshal(first)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = json.Unmarshal(js, second)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	reJS, err := json.Marshal(second)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(js, reJS) != 0 {
		t.Errorf("the transformed javascript is different.\n%s\n%s", js, reJS)
		return
	}

	if !reflect.DeepEqual(first, second) {
		t.Errorf("the instance conversion failed")
		return
	}
}
