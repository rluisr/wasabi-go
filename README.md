wasabi-go
==========
![golangci-lint](https://github.com/rluisr/wasabi-go/workflows/golangci-lint/badge.svg)

Client for Wasabi Wallet.  
This is under the work in progress.

```go
package main

import "github.com/rluisr/wasabi-go"

func main() {
    client, err = wasabi.New("http://localhost:37128")
    err = client.SelectWallet("Wallet0")

    wallet, err := client.GetWalletInfo()
}
```

Supported method
----------------
- [ ] get status
- [x] create wallet
- [ ] list unspent coins
- [x] get wallet info
- [x] get new address
- [ ] send
- [ ] get history
- [ ] list keys
- [ ] enqueue
- [ ] dequeue
- [ ] stop