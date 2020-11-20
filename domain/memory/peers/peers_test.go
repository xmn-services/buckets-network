package peers

import (
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

func TestPeers_Success(t *testing.T) {
	lst := []peer.Peer{
		peer.CreatePeerWithClear("12.47.89.78", 80),
		peer.CreatePeerWithOnion("asfsdfdfgtebhndf", 443),
	}

	ins, err := NewBuilder().Create().WithPeers(lst).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil")
		return
	}

	repository, service := CreateRepositoryServiceForTests()
	err = service.Save(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil")
		return
	}

	retIns, err := repository.Retrieve()
	if err != nil {
		t.Errorf("the error was expected to be nil")
		return
	}

	TestCompare(t, ins, retIns)
}
