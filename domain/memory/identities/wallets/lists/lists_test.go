package lists

import (
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"
)

func TestLists_Success(t *testing.T) {
	first := list.CreateListForTests()
	second := list.CreateListForTests()
	third := list.CreateListForTests()

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

	err = ins.Delete(first.Hash())
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
	retIns, err := adapter.ToLists(js)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)

}
