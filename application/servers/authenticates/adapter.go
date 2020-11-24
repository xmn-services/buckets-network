package authenticates

import (
	b64 "encoding/base64"
	"encoding/json"
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

// Base64ToAuthenticate converts a base64 encoded string to an Authenticate instance
func (app *adapter) Base64ToAuthenticate(encoded string) (Authenticate, error) {
	decoded, err := b64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	mp := map[string]string{}
	err = json.Unmarshal(decoded, &mp)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create()
	if name, ok := mp["name"]; ok {
		builder.WithName(name)
	}

	if password, ok := mp["password"]; ok {
		builder.WithPassword(password)
	}

	if seed, ok := mp["seed"]; ok {
		builder.WithSeed(seed)
	}

	return builder.Now()
}

// AuthenticateToBase64 converts an Authenticate instance to base64 encoded string
func (app *adapter) AuthenticateToBase64(auth Authenticate) (string, error) {
	mp := map[string]string{
		"name":     auth.Name(),
		"password": auth.Password(),
		"seed":     auth.Seed(),
	}

	js, err := json.Marshal(mp)
	if err != nil {
		return "", nil
	}

	return b64.StdEncoding.EncodeToString(js), nil
}
