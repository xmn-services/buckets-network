package peer

import (
	"fmt"
	"testing"
)

func TestPeer_withClear_Success(t *testing.T) {
	host := "78.98.32.34"
	port := uint(8080)
	peer := CreatePeerWithClear(host, port)

	if peer.Port() != port {
		t.Errorf("the port was expected to be %d, %d returned", port, peer.Port())
		return
	}

	if peer.Host() != host {
		t.Errorf("the host was expected to be %s, %s returned", host, peer.Host())
		return
	}

	if !peer.IsClear() {
		t.Errorf("the peer was expected to be a clear peer")
		return
	}

	if peer.IsOnion() {
		t.Errorf("the peer was expected to NOT be an onion peer")
		return
	}

	str := fmt.Sprintf("http://%s:%d", host, port)
	if peer.String() != str {
		t.Errorf("the peer string is invalid, expected: %s, returned %s", str, peer.String())
		return
	}

	adapter := NewAdapter()
	urlValues := adapter.PeerToURLValues(peer)
	retPeer, err := adapter.URLValuesToPeer(urlValues)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retPeer.String() != str {
		t.Errorf("the peer string is invalid, expected: %s, returned %s", str, retPeer.String())
		return
	}
}

func TestPeer_withOnion_Success(t *testing.T) {
	host := "asdfasdfasdf"
	port := uint(443)
	peer := CreatePeerWithOnion(host, port)

	if peer.Port() != port {
		t.Errorf("the port was expected to be %d, %d returned", port, peer.Port())
		return
	}

	if peer.Host() != host {
		t.Errorf("the host was expected to be %s, %s returned", host, peer.Host())
		return
	}

	if peer.IsClear() {
		t.Errorf("the peer was expected to NOT be a clear peer")
		return
	}

	if !peer.IsOnion() {
		t.Errorf("the peer was expected to be an onion peer")
		return
	}

	str := fmt.Sprintf("onion://%s:%d", host, port)
	if peer.String() != str {
		t.Errorf("the peer string is invalid, expected: %s, returned %s", str, peer.String())
		return
	}

	adapter := NewAdapter()
	urlValues := adapter.PeerToURLValues(peer)
	retPeer, err := adapter.URLValuesToPeer(urlValues)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retPeer.String() != str {
		t.Errorf("the peer string is invalid, expected: %s, returned %s", str, retPeer.String())
		return
	}
}
