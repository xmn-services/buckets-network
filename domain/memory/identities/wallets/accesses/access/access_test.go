package access

import (
	"reflect"
	"testing"

	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestAccess_Success(t *testing.T) {
	bucket, _ := hash.NewAdapter().FromBytes([]byte("bucket hash"))
	ins, key := CreateAccessForTests(*bucket)

	retBucket := ins.Bucket()
	if !bucket.Compare(retBucket) {
		t.Errorf("the returned bucket is invalid")
		return
	}

	retKey := ins.Key()
	if !reflect.DeepEqual(key, retKey) {
		t.Errorf("the returned publicKey is invalid")
		return
	}

	adapter := NewAdapter()
	js := adapter.ToJSON(ins)
	retIns, err := adapter.ToAccess(js)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)

}
