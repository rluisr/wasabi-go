package wasabi

import "encoding/json"

type Address struct {
	Address   string   `json:"address"`
	KeyPath   string   `json:"keyPath"`
	Label     []string `json:"label"`
	PublicKey string   `json:"publicKey"`
	P2wpkh    string   `json:"p2wpkh"`
}

func (c *Client) GetNewAddress(name string) (*Address, error) {
	resp, err := c.request(Request{
		JSONRPC: JSONRPC,
		ID:      "1",
		Method:  "getnewaddress",
		Params:  []string{name},
	})
	if err != nil {
		return nil, err
	}

	var address Address
	err = json.Unmarshal(resp.Result, &address)

	return &address, err
}
