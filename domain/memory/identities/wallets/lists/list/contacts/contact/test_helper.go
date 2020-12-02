package contact

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact/accesses"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
)

// CreateContactForTests creates a contact for tests
func CreateContactForTests() Contact {
	pk, _ := encryption.NewFactory(4096).Create()
	key := pk.Public()

	name := "Roger Cyr"
	description := "this is a description"
	accesses, _ := accesses.NewFactory().Create()

	ins, err := NewBuilder().Create().WithKey(key).WithName(name).WithDescription(description).WithAccess(accesses).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// TestCompare compare two contact instances
func TestCompare(t *testing.T, first Contact, second Contact) {
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
