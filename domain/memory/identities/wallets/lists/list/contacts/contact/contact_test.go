package contact

import (
	"testing"
)

func TestContact_Success(t *testing.T) {
	ins := CreateContactForTests()
	adapter := NewAdapter()
	js := adapter.ToJSON(ins)
	retIns, err := adapter.ToContact(js)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)
}
