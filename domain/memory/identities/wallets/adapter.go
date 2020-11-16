package wallets

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts Wallet instance to JSON
func (app *adapter) ToJSON(ins Wallet) *JSONWallet {
	return createJSONWalletFromWallet(ins)
}

// ToWallet converts JSON to Wallet instance
func (app *adapter) ToWallet(js *JSONWallet) (Wallet, error) {
	return createWalletFromJSON(js)
}
