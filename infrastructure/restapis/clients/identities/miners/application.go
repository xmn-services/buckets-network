package miners

import "github.com/xmn-services/buckets-network/application/commands/identities/miners"

type application struct {
}

func createApplication() miners.Application {
	out := application{}
	return &out
}

// Test executes a test on the miner application
func (app *application) Test(difficulty uint) (string, error) {
	return "", nil
}

// Block mines a block
func (app *application) Block(blockHashStr string) (string, error) {
	return "", nil
}

// Link mines a link
func (app *application) Link(linkHashStr string) (string, error) {
	return "", nil
}
