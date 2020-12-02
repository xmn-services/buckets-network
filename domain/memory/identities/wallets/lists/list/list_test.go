package list

import (
	"testing"
)

func TestList_Success(t *testing.T) {
	ins := CreateListForTests()
	adapter := NewAdapter()
	js := adapter.ToJSON(ins)
	retIns, err := adapter.ToList(js)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)
}
