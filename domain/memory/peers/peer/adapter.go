package peer

import (
	"net/url"
	"strconv"
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

// URLValuesToPeer converts url values to peer
func (app *adapter) URLValuesToPeer(values url.Values) (Peer, error) {
	str := values.Get("peer")
	return app.StringToPeer(str)
}

// PeerToURLValues converts a Peer instance to url values
func (app *adapter) PeerToURLValues(peer Peer) url.Values {
	values := url.Values{}
	values.Add("peer", peer.String())
	return values
}

// StringToPeer converts a string to a peer instance
func (app *adapter) StringToPeer(str string) (Peer, error) {
	url, err := url.Parse(str)
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(url.Port())
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithHost(url.Hostname()).WithPort(uint(port))
	switch url.Scheme {
	case "http":
		builder.IsClear()
		break
	case "onion":
		builder.IsOnion()
		break
	}

	return builder.Now()
}
