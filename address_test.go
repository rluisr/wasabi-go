package wasabi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetNewAddress(t *testing.T) {
	setup()

	address, err := client.GetNewAddress("test")
	assert.NoError(t, err)
	assert.NotEmpty(t, address)
}
