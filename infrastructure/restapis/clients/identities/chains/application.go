package chains

import "github.com/xmn-services/buckets-network/application/commands/identities/chains"

type application struct {
}

func createApplication() chains.Application {
	out := application{}
	return &out
}

// Init initializes a chain application
func (app *application) Init(
	miningValue uint8,
	baseDifficulty uint,
	increasePerBucket float64,
	linkDifficulty uint,
	rootAdditionalBuckets uint,
	headAdditionalBuckets uint,
) error {
	return nil
}

// Block mines a block on the chain
func (app *application) Block(additional uint) error {
	return nil
}

// Link mines a link on the chain
func (app *application) Link(additional uint) error {
	return nil
}
