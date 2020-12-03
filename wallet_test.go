package wasabi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_SelectWallet(t *testing.T) {
	setup()

	err := client.SelectWallet("Wallet0")
	assert.NoError(t, err)
}

func TestClient_GetWalletInfo(t *testing.T) {
	setup()

	wallet, err := client.GetWalletInfo()
	assert.NoError(t, err)
	assert.NotEmpty(t, wallet)
}

//func TestClient_CreateWallet(t *testing.T) {
//	setup()
//
//	recoveryKeys, err := client.CreateWallet("test", "test")
//	assert.NoError(t, err)
//	assert.NotEmpty(t, recoveryKeys)
//}
