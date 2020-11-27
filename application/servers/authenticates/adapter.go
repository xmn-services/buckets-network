package authenticates

import (
	b64 "encoding/base64"
	"encoding/json"
	"net/url"
)

type adapter struct {
	builder Builder
}

func createAdapter(
	builder Builder,
) Adapter {
	out := adapter{
		builder: builder,
	}

	return &out
}

// URLValuesToAuthenticate converts an url values to Authenticate instance
func (app *adapter) URLValuesToAuthenticate(urlValues url.Values) (Authenticate, error) {
	builder := app.builder.Create()
	name := urlValues.Get("name")
	if name != "" {
		builder.WithName(name)
	}

	password := urlValues.Get("password")
	if password != "" {
		builder.WithPassword(password)
	}

	seed := urlValues.Get("seed")
	if seed != "" {
		builder.WithSeed(seed)
	}

	return builder.Now()
}

// Base64ToAuthenticate converts a base64 encoded string to an Authenticate instance
func (app *adapter) Base64ToAuthenticate(encoded string) (Authenticate, error) {
	decoded, err := b64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	urlValues := new(url.Values)
	err = json.Unmarshal(decoded, urlValues)
	if err != nil {
		return nil, err
	}

	return app.URLValuesToAuthenticate(*urlValues)
}

// AuthenticateToURLValues converts an authenticate to urlValues instance
func (app *adapter) AuthenticateToURLValues(auth Authenticate) url.Values {
	urlValues := url.Values{}
	urlValues.Add("name", auth.Name())
	urlValues.Add("password", auth.Password())
	urlValues.Add("seed", auth.Seed())
	return urlValues
}

// AuthenticateToBase64 converts an Authenticate instance to base64 encoded string
func (app *adapter) AuthenticateToBase64(auth Authenticate) (string, error) {
	values := app.AuthenticateToURLValues(auth)
	js, err := json.Marshal(values)
	if err != nil {
		return "", nil
	}

	return b64.StdEncoding.EncodeToString(js), nil
}
