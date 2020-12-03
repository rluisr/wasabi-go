wasabi-go
==========
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
- [ ] create wallet
- [ ] list unspent coins
- [x] get wallet info
- [ ] get new address
- [ ] send
- [ ] get history
- [ ] list keys
- [ ] enqueue
- [ ] dequeue
- [ ] stop