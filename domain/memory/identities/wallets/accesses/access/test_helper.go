package access

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// CreateAccessForTests creates an access instance for tests
func CreateAccessForTests(bucket hash.Hash) (Access, public.Key) {
	pk, _ := encryption.NewFactory(4096).Create()
	key := pk.Public()

	ins, err := NewBuilder().Create().WithBucket(bucket).WithKey(key).Now()
	if err != nil {
		panic(err)
	}

	return ins, key
}

// TestCompare compare two access instances
func TestCompare(t *testing.T, first Access, second Access) {
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
