package examples

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	"github.com/sonirico/go-hyperliquid"
)

func TestWebsocket(t *testing.T) {
	ws := hyperliquid.NewWebsocketClient(hyperliquid.MainnetAPIURL)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Log("Connecting to websocket")
	if err := ws.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer ws.Close()

	t.Log("Connected to websocket")

	// Test trades subscription
	t.Run("trades subscription", func(t *testing.T) {
		wg := new(sync.WaitGroup)
		wg.Add(1)
		received := make(chan struct{})

		tradesSub := hyperliquid.Subscription{
			Type: "trades",
			Coin: "SOL",
		}

		_, err := ws.Subscribe(tradesSub, func(msg hyperliquid.WSMessage) {
			trades := []hyperliquid.Trade{}
			if err := json.Unmarshal(msg.Data, &trades); err != nil {
				t.Fatalf("Failed to unmarshal trades: %v", err)
			}

			t.Logf("Received trade: %+v", trades)
			wg.Done()
		})

		if err != nil {
			t.Fatalf("Failed to subscribe to trades: %v", err)
		}

		go func() {
			wg.Wait()
			close(received)
		}()

		select {
		case <-received:
			// Test passed
		case <-ctx.Done():
			t.Fatal("Timeout waiting for trades messages")
		}
	})

	// Test L2 book subscription
	t.Run("l2book subscription", func(t *testing.T) {
		wg := new(sync.WaitGroup)
		wg.Add(1)
		received := make(chan struct{})

		l2Sub := hyperliquid.Subscription{
			Type: "l2Book",
			Coin: "BTC",
		}

		_, err := ws.Subscribe(l2Sub, func(msg hyperliquid.WSMessage) {
			l2Update := hyperliquid.L2Book{}
			if err := json.Unmarshal(msg.Data, &l2Update); err != nil {
				t.Fatalf("Failed to unmarshal L2 update: %v", err)
			}
			t.Logf("Received L2 update: %+v", l2Update)
			wg.Done()
		})

		if err != nil {
			t.Fatalf("Failed to subscribe to L2 book: %v", err)
		}

		go func() {
			wg.Wait()
			close(received)
		}()

		select {
		case <-received:
			// Test passed
		case <-ctx.Done():
			t.Fatal("Timeout waiting for L2 book messages")
		}
	})
}
