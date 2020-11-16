package wallets

import (
	"testing"
)

func TestWallet_Success(t *testing.T) {
	ins, err := NewFactory().Create()
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	adapter := NewAdapter()
	jsIns := adapter.ToJSON(ins)
	retIns, err := adapter.ToWallet(jsIns)
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)

}
