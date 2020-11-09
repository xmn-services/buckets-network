package addresses

import "time"

// JSONAddress represents a json address
type JSONAddress struct {
	Sender     string    `json:"sender"`
	Recipients []string  `json:"recipients"`
	CreatedOn  time.Time `json:"created_on"`
}

func createJSONAddressFromAddress(address Address) *JSONAddress {
	sender := ""
	if address.HasSender() {
		sender = address.Sender().String()
	}

	recipients := []string{}
	if address.HasRecipients() {
		lst := address.Recipients()
		for _, oneRecipient := range lst {
			recipients = append(recipients, oneRecipient.String())
		}
	}

	createdOn := address.CreatedOn()
	return createJSONAddress(sender, recipients, createdOn)
}

func createJSONAddress(
	sender string,
	recipients []string,
	createdOn time.Time,
) *JSONAddress {
	out := JSONAddress{
		Sender:     sender,
		Recipients: recipients,
		CreatedOn:  createdOn,
	}

	return &out
}
