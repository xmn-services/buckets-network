package peers

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

type peers struct {
	peers []peer.Peer
}

func createPeersFromStrings(lst []string) (Peers, error) {
	peers := []peer.Peer{}
	peerAdapter := peer.NewAdapter()
	for _, oneURLStr := range lst {
		peer, err := peerAdapter.StringToPeer(oneURLStr)
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
