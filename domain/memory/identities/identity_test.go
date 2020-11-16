package identities

import (
	"testing"
)

func TestWallet_Success(t *testing.T) {
	seed := "this is my seed"
	name := "my_name"
	root := "/my/root"
	pass := "this is my password"

	basePath := "./test_files"
	extension := "json"

	ins, err := NewBuilder().Create().WithSeed(seed).WithName(name).WithRoot(root).Now()
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	if ins.Seed() != seed {
		t.Errorf("the seed is invalid: expected: %s, returned: %s", seed, ins.Seed())
		return
	}

	if ins.Name() != name {
		t.Errorf("the name is invalid: expected: %s, returned: %s", name, ins.Name())
		return
	}

	if ins.Root() != root {
		t.Errorf("the root is invalid: expected: %s, returned: %s", root, ins.Root())
		return
	}

	repository := NewRepository(basePath, extension)
	service := NewService(basePath, extension)

	err = service.Insert(ins, pass)
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	retIns, err := repository.Retrieve(name, pass, seed)
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	err = service.Delete(ins, pass)
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	_, err = repository.Retrieve(name, pass, seed)
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}

	TestCompare(t, ins, retIns)

}
