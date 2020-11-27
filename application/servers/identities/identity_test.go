package identities

import (
	"reflect"
	"testing"

	"github.com/xmn-services/buckets-network/application/servers/authenticates"
)

func TestIdentity_Success(t *testing.T) {
	name := "my_name"
	password := "asfsfasfq234235"
	seed := "this is my seed"
	root := "/my/root/path"
	auth, err := authenticates.NewBuilder().Create().WithName(name).WithPassword(password).WithSeed(seed).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins, err := NewBuilder().Create().WithAuthenticate(auth).WithRoot(root).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewAdapter()
	urlValues := adapter.IdentityToURLValues(ins)
	retIdentity, err := adapter.URLValuesToIdentity(urlValues)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(ins, retIdentity) {
		t.Errorf("the identity is invalid")
		return
	}

}
