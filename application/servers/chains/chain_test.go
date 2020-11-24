package chains

import (
	"testing"
)

func TestChain_Success(t *testing.T) {
	miningValue := uint8(0)
	baseDiff := uint(1)
	incrPerBucket := float64(0.0005)
	likDiff := uint(2)
	rootAddBuckets := uint(20)
	headAddBuckets := uint(30)
	chain, err := NewBuilder().
		Create().
		WithMiningValue(miningValue).
		WithBaseDifficulty(baseDiff).
		WithIncreasePerBucket(incrPerBucket).
		WithLinkDifficulty(likDiff).
		WithRootAdditionalBuckets(rootAddBuckets).
		WithHeadAdditionalBuckets(headAddBuckets).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewAdapter()
	values := adapter.ChainToURLValues(chain)
	retChain, err := adapter.URLValuesToChain(values)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if chain.MiningValue() != retChain.MiningValue() {
		t.Errorf("the miningValue was expected to be %d, %d returned", chain.MiningValue(), retChain.MiningValue())
		return
	}

	if chain.BaseDifficulty() != retChain.BaseDifficulty() {
		t.Errorf("the baseDifficulty was expected to be %d, %d returned", chain.BaseDifficulty(), retChain.BaseDifficulty())
		return
	}

	if chain.IncreasePerBucket() != retChain.IncreasePerBucket() {
		t.Errorf("the increasePerBucket was expected to be %f, %f returned", chain.IncreasePerBucket(), retChain.IncreasePerBucket())
		return
	}

	if chain.LinkDifficulty() != retChain.LinkDifficulty() {
		t.Errorf("the linkDifficulty was expected to be %d, %d returned", chain.LinkDifficulty(), retChain.LinkDifficulty())
		return
	}

	if chain.RootAdditionalBuckets() != retChain.RootAdditionalBuckets() {
		t.Errorf("the rootAdditionalBuckets was expected to be %d, %d returned", chain.RootAdditionalBuckets(), retChain.RootAdditionalBuckets())
		return
	}

	if chain.HeadAdditionalBuckets() != retChain.HeadAdditionalBuckets() {
		t.Errorf("the headAdditionalBuckets was expected to be %d, %d returned", chain.HeadAdditionalBuckets(), retChain.HeadAdditionalBuckets())
		return
	}
}
