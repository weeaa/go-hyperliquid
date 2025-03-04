package examples

import (
	"testing"
)

func TestUpdateLeverage(t *testing.T) {
	exchange := getTestExchange(t)

	leverage := 5 // 5x leverage
	coin := "BTC"

	resp, err := exchange.UpdateLeverage(coin, leverage)
	if err != nil {
		t.Fatalf("Failed to update leverage: %v", err)
	}

	t.Logf("Update leverage response: %+v", resp)
}

func TestUpdateIsolatedMargin(t *testing.T) {
	exchange := getTestExchange(t)

	amount := 1000.0 // Amount in USD
	coin := "BTC"

	resp, err := exchange.UpdateIsolatedMargin(coin, amount)
	if err != nil {
		t.Fatalf("Failed to update isolated margin: %v", err)
	}

	t.Logf("Update isolated margin response: %+v", resp)
}
