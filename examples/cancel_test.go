package examples

import (
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/weeaa/go-hyperliquid"
)

func TestCancelOrder(t *testing.T) {
	exchange := getTestExchange(t)

	// First place an order to cancel
	orderReq := hyperliquid.OrderRequest{
		Coin:    "BTC",
		IsBuy:   true,
		Size:    0.1,
		LimitPx: 40000.0,
		OrderType: hyperliquid.OrderType{
			Limit: &hyperliquid.LimitOrderType{
				Tif: "Gtc",
			},
		},
	}

	resp, err := exchange.Order(orderReq, nil)
	if err != nil {
		t.Fatalf("Failed to place order: %v", err)
	}

	// Extract order ID from response
	orderID := resp.Oid

	// Cancel the order
	cancelResp, err := exchange.Cancel("BTC", orderID)
	if err != nil {
		t.Fatalf("Failed to cancel order: %v", err)
	}

	t.Logf("Cancel response: %+v", cancelResp)
}

func TestCancelByCloid(t *testing.T) {
	exchange := getTestExchange(t)

	// Generate a random cloid
	cloidBytes := make([]byte, 16)
	if _, err := rand.Read(cloidBytes); err != nil {
		t.Fatalf("Failed to generate random cloid: %v", err)
	}
	cloid := "0x" + hex.EncodeToString(cloidBytes)

	// Place an order with cloid
	orderReq := hyperliquid.OrderRequest{
		Coin:    "BTC",
		IsBuy:   true,
		Size:    0.1,
		LimitPx: 40000.0,
		OrderType: hyperliquid.OrderType{
			Limit: &hyperliquid.LimitOrderType{
				Tif: "Gtc",
			},
		},
		Cloid: &cloid,
	}

	_, err := exchange.Order(orderReq, nil)
	if err != nil {
		t.Fatalf("Failed to place order: %v", err)
	}

	// Cancel by cloid
	cancelResp, err := exchange.CancelByCloid("BTC", cloid)
	if err != nil {
		t.Fatalf("Failed to cancel order by cloid: %v", err)
	}

	t.Logf("Cancel by cloid response: %+v", cancelResp)
}
