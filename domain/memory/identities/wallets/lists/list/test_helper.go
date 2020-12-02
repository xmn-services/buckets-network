package list

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts"
)

// CreateListForTests creates a list for tests
func CreateListForTests() List {
	name := "Roger Cyr"
	description := fmt.Sprintf("this is a description: %s", time.Now().UTC().String())
	contacts, _ := contacts.NewFactory().Create()

	ins, err := NewBuilder().Create().WithName(name).WithDescription(description).WithContacts(contacts).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// TestCompare compare two list instances
func TestCompare(t *testing.T, first List, second List) {
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
