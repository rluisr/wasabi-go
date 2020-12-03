package wasabi

import (
	"encoding/json"
)

type Wallet struct {
	WalletName               string `json:"walletName"`
	WalletFile               string `json:"walletFile"`
	State                    string `json:"state"`
	ExtendedAccountPublicKey string `json:"extendedAccountPublicKey"`
	ExtendedAccountZpub      string `json:"extendedAccountZpub"`
	AccountKeyPath           string `json:"accountKeyPath"`
	MasterKeyFingerprint     string `json:"masterKeyFingerprint"`
	Balance                  uint   `json:"balance"`
}

func (c *Client) SelectWallet(name string) error {
	_, err := c.request(Request{
		JSONRPC: JSONRPC,
		Method:  "selectwallet",
		Params:  []string{name},
	})

	return err
}

func (c *Client) GetWalletInfo() (*Wallet, error) {
	resp, err := c.request(Request{
		JSONRPC: JSONRPC,
		ID:      "1",
		Method:  "getwalletinfo",
	})
	if err != nil {
		return nil, err
	}

	var wallet Wallet
	err = json.Unmarshal(resp.Result, &wallet)

	return &wallet, err
}
