package wasabi

const JSONRPC = "2.0"

type Client struct {
	URL string
}

func newClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

func New(url string) (*Client, error) {
	client := newClient(url)
	err := client.verifyClient()

	return client, err
}

func (c *Client) verifyClient() error {
	_, err := c.request(Request{
		JSONRPC: JSONRPC,
		ID:      "1",
		Method:  "getstatus",
	})

	return err
}
