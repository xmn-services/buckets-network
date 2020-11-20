package peer

// CreatePeerWithClear creates a new peer instance with clear
func CreatePeerWithClear(host string, port uint) Peer {
	ins, err := NewBuilder().Create().IsClear().WithHost(host).WithPort(port).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreatePeerWithOnion creates a new peer instance with onion
func CreatePeerWithOnion(host string, port uint) Peer {
	ins, err := NewBuilder().Create().IsOnion().WithHost(host).WithPort(port).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
