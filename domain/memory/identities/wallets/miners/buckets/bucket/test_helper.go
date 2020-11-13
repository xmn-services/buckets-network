package bucket

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
)

// CreateBucketForTestsWithoutParams creates a bucket for tests without params
func CreateBucketForTestsWithoutParams() Bucket {
	absolutePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	bucketIns := buckets.CreateBucketForTestsWithoutParams()
	pk, err := encryption.NewFactory(4096).Create()
	if err != nil {
		panic(err)
	}

	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithBucket(bucketIns).WithAbsolutePath(absolutePath).WithPrivateKey(pk).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// TestCompare compare two bucket instances
func TestCompare(t *testing.T, first Bucket, second Bucket) {
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

	if !first.Hash().Compare(second.Hash()) {
		t.Errorf("the instance conversion failed")
		return
	}
}
