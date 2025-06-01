package examples

import (
	"crypto/ecdsa"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/weeaa/go-hyperliquid"
)

var (
	testPrivateKey *ecdsa.PrivateKey
	testExchange   *hyperliquid.Exchange
)

func TestMain(m *testing.M) {
	// Setup test environment
	var err error
	privKeyHex := os.Getenv("HL_PRIVATE_KEY")
	if privKeyHex == "" {
		// Use a test private key if none provided
		privKeyHex = "ff"
	}

	testPrivateKey, err = crypto.HexToECDSA(privKeyHex)
	if err != nil {
		panic("failed to parse private key: " + err.Error())
	}

	// Initialize test exchange
	testExchange = hyperliquid.NewExchange(
		testPrivateKey,
		hyperliquid.MainnetAPIURL,
		nil,
		os.Getenv("HL_VAULT_ADDRESS"),
		crypto.PubkeyToAddress(testPrivateKey.PublicKey).Hex(),
		nil,
	)

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func getTestExchange(t *testing.T) *hyperliquid.Exchange {
	t.Helper()
	if testExchange == nil {
		t.Fatal("test exchange not initialized")
	}
	return testExchange
}

func skipIfNoPrivateKey(t *testing.T) {
	t.Helper()
	if os.Getenv("HL_PRIVATE_KEY") == "" {
		t.Skip("Skipping test: no private key provided")
	}
}
