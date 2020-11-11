package buckets

import (
	"encoding/json"
	"testing"
)

func TestFile_Success(t *testing.T) {
	bucketIns := CreateBucketForTestsWithoutParams()

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(bucketIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// save the file:
	err = service.Save(bucketIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBucket, err := repository.Retrieve(bucketIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bucketIns.Hash().Compare(retBucket.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(bucketIns)
	if err != nil {
		t.Errorf("the File instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(bucket)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a File instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, bucketIns, retBucket)
}
