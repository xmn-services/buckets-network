package bucket

import (
	"testing"
)

func TestBucket_Success(t *testing.T) {
	ins := CreateBucketForTestsWithoutParams()

	adapter := NewAdapter()
	jsIns := adapter.ToJSON(ins)
	retIns, err := adapter.ToBucket(jsIns)
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)
}
