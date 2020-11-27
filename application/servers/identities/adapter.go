package identities

import (
	"net/url"

	"github.com/xmn-services/buckets-network/application/servers/authenticates"
)

type adapter struct {
	authAdapter authenticates.Adapter
	builder     Builder
}

func createAdapter(
	authAdapter authenticates.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		authAdapter: authAdapter,
		builder:     builder,
	}

	return &out
}

// URLValuesToIdentity converts an url values to an Identity instance
func (app *adapter) URLValuesToIdentity(urlValues url.Values) (Identity, error) {
	auth, err := app.authAdapter.URLValuesToAuthenticate(urlValues)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create()
	root := urlValues.Get("root")
	if root != "" {
		builder.WithRoot(root)
	}

	return builder.WithAuthenticate(auth).Now()
}

// IdentityToURLValues converts an Identity instance to url values
func (app *adapter) IdentityToURLValues(identity Identity) url.Values {
	values := app.authAdapter.AuthenticateToURLValues(identity.Authenticate())
	values.Add("root", identity.Root())
	return values
}
