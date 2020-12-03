package wasabi

var (
	client *Client
)

func setup() {
	var err error

	client, err = New("http://localhost:37128")
	if err != nil {
		panic(err)
	}
}
