package permanents

import (
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
)

func TestBuckets_withoutBuckets_Success(t *testing.T) {
	ins, err := NewFactory().Create()
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	adapter := NewAdapter()
	jsIns := adapter.ToJSON(ins)
	retIns, err := adapter.ToBuckets(jsIns)
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)
}

func TestBuckets_addBucket_Success(t *testing.T) {
	ins, err := NewFactory().Create()
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	bucket := bucket.CreateBucketForTestsWithoutParams()
	err = ins.Add(bucket)
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	adapter := NewAdapter()
	jsIns := adapter.ToJSON(ins)
	retIns, err := adapter.ToBuckets(jsIns)
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)
}
