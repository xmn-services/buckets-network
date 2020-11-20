package peers

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

type peers struct {
	peers []peer.Peer
}

func createPeersFromStrings(lst []string) (Peers, error) {
	peers := []peer.Peer{}
	peerBuilder := peer.NewBuilder()
	for _, oneURLStr := range lst {
		url, err := url.Parse(oneURLStr)
		if err != nil {
			return nil, err
		}

		port, err := strconv.Atoi(url.Port())
		if err != nil {
			return nil, err
		}

		builder := peerBuilder.Create().WithHost(url.Hostname()).WithPort(uint(port))
		switch url.Scheme {
		case "https":
			builder.IsClear()
			break
		case "onion":
			builder.IsOnion()
			break
		}

		peer, err := builder.Now()
		if err != nil {
			return nil, err
		}

		peers = append(peers, peer)
	}

	return createPeers(peers), nil
}

func createPeers(
	lst []peer.Peer,
) Peers {
	out := peers{
		peers: lst,
	}

	return &out
}

// All returns the peers
func (obj *peers) All() []peer.Peer {
	return obj.peers
}

// MarshalJSON converts the instance to JSON
func (obj *peers) MarshalJSON() ([]byte, error) {
	lst := []string{}
	for _, onePeer := range obj.peers {
		lst = append(lst, onePeer.String())
	}

	return json.Marshal(lst)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *peers) UnmarshalJSON(data []byte) error {
	ins := new([]string)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createPeersFromStrings(*ins)
	if err != nil {
		return err
	}

	insPeers := pr.(*peers)
	obj.peers = insPeers.peers
	return nil
}
