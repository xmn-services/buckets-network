package authenticates

import (
	"testing"
)

func TestAuthenticate_Success(t *testing.T) {
	name := "my_name"
	password := "asfsfasfq234235"
	seed := "this is my seed"
	ins, err := NewBuilder().Create().WithName(name).WithPassword(password).WithSeed(seed).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewAdapter()
	encoded, err := adapter.AuthenticateToBase64(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	decoded, err := adapter.Base64ToAuthenticate(encoded)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins.Name() != decoded.Name() {
		t.Errorf("the name was expected to be %s, %s returned", ins.Name(), decoded.Name())
		return
	}

	if ins.Seed() != decoded.Seed() {
		t.Errorf("the seed was expected to be %s, %s returned", ins.Seed(), decoded.Seed())
		return
	}

	if ins.Password() != decoded.Password() {
		t.Errorf("the password was expected to be %s, %s returned", ins.Password(), decoded.Password())
		return
	}
}
