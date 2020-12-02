package accesses

import (
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses/access"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestAccesses_Success(t *testing.T) {
	firstBucket, _ := hash.NewAdapter().FromBytes([]byte("first bucket hash"))
	first, _ := access.CreateAccessForTests(*firstBucket)

	secondBucket, _ := hash.NewAdapter().FromBytes([]byte("second bucket hash"))
	second, _ := access.CreateAccessForTests(*secondBucket)

	thirdBucket, _ := hash.NewAdapter().FromBytes([]byte("third bucket hash"))
	third, _ := access.CreateAccessForTests(*thirdBucket)

	ins, err := NewFactory().Create()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = ins.Add(first)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = ins.Add(second)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = ins.Add(third)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = ins.Delete(first.Bucket())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retLst := ins.All()
	if len(retLst) != 2 {
		t.Errorf("%d instances were expected, %d returned", 2, len(retLst))
		return
	}

	adapter := NewAdapter()
	js := adapter.ToJSON(ins)
	retIns, err := adapter.ToAccesses(js)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)

}
