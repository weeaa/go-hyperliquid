package examples

import (
	"testing"

	"github.com/sonirico/go-hyperliquid"
)

func TestOrder(t *testing.T) {
	skipIfNoPrivateKey(t)
	exchange := getTestExchange(t)

	tests := []struct {
		name string
		req  hyperliquid.OrderRequest
	}{
		{
			name: "limit buy order",
			req: hyperliquid.OrderRequest{
				Coin:    "BTC",
				IsBuy:   true,
				Size:    0.001, // Smaller size for testing
				LimitPx: 40000.0,
				OrderType: hyperliquid.OrderType{
					Limit: &hyperliquid.LimitOrderType{
						Tif: "Gtc",
					},
				},
			},
		},
		{
			name: "market sell order",
			req: hyperliquid.OrderRequest{
				Coin:    "ETH",
				IsBuy:   false,
				Size:    0.01,
				LimitPx: 2000.0,
				OrderType: hyperliquid.OrderType{
					Limit: &hyperliquid.LimitOrderType{
						Tif: "Ioc",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := exchange.Order(tt.req, nil)
			if err != nil {
				t.Fatalf("Order failed: %v", err)
			}
			t.Logf("Order response: %+v", resp)
		})
	}
}
