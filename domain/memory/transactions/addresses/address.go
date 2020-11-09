package addresses

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type address struct {
	immutable  entities.Immutable
	sender     *hash.Hash
	recipients []hash.Hash
}

func createAddressFromJSON(js *JSONAddress) (Address, error) {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder().Create().CreatedOn(js.CreatedOn)
	if js.Sender != "" {
		sender, err := hashAdapter.FromString(js.Sender)
		if err != nil {
			return nil, err
		}

		builder.WithSender(*sender)
	}

	if len(js.Recipients) > 0 {
		recipients := []hash.Hash{}
		for _, oneRecipient := range js.Recipients {
			ins, err := hashAdapter.FromString(oneRecipient)
			if err != nil {
				return nil, err
			}

			recipients = append(recipients, *ins)
		}

		builder.WithRecipients(recipients)
	}

	return builder.Now()
}

func createAddressWithSender(
	immutable entities.Immutable,
	sender *hash.Hash,
) Address {
	return createAddressInternally(immutable, sender, nil)
}

func createAddressWithRecipients(
	immutable entities.Immutable,
	recipients []hash.Hash,
) Address {
	return createAddressInternally(immutable, nil, recipients)
}

func createAddressWithSenderAndRecipients(
	immutable entities.Immutable,
	sender *hash.Hash,
	recipients []hash.Hash,
) Address {
	return createAddressInternally(immutable, sender, recipients)
}

func createAddressInternally(
	immutable entities.Immutable,
	sender *hash.Hash,
	recipients []hash.Hash,
) Address {
	out := address{
		immutable:  immutable,
		sender:     sender,
		recipients: recipients,
	}

	return &out
}

// Hash returns the hash
func (obj *address) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// CreatedOn returns the creation time
func (obj *address) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasSender returns true if there is a sender, false otherwise
func (obj *address) HasSender() bool {
	return obj.sender != nil
}

// Sender returns the sender, if any
func (obj *address) Sender() *hash.Hash {
	return obj.sender
}

// HasRecipients returns true if there is recipients, false otherwise
func (obj *address) HasRecipients() bool {
	return obj.recipients != nil
}

// Recipients returns the recipients, if any
func (obj *address) Recipients() []hash.Hash {
	return obj.recipients
}

// MarshalJSON converts the instance to JSON
func (obj *address) MarshalJSON() ([]byte, error) {
	ins := createJSONAddressFromAddress(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *address) UnmarshalJSON(data []byte) error {
	ins := new(JSONAddress)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createAddressFromJSON(ins)
	if err != nil {
		return err
	}

	insAddress := pr.(*address)
	obj.immutable = insAddress.immutable
	obj.sender = insAddress.sender
	obj.recipients = insAddress.recipients
	return nil
}
