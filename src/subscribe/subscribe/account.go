package subscribe

import (
	"context"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func AccountSubscribe(url, pub string) error {
	fmt.Println("starting account subscription...")

	// Connect to Solana WebSocket endpoint
	client, err := ws.Connect(context.Background(), url)
	if err != nil {
		return fmt.Errorf("webSocket connection failed: %w", err)
	}
	defer client.Close()

	// Replace with the account public key you want to watch
	pubkey, err := solana.PublicKeyFromBase58(pub)
	if err != nil {
		return fmt.Errorf("invalid public key: %w", err)
	}

	// Subscribe to account changes
	sub, err := client.AccountSubscribe(pubkey, rpc.CommitmentConfirmed)
	if err != nil {
		return fmt.Errorf("account subscribe error: %w", err)
	}
	defer sub.Unsubscribe()

	// set outer timeout (e.g. total listen time)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Listen for account updates
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("subscription timeout: %w", ctx.Err())
		default:
			msg, err := sub.Recv(context.Background()) // no timeout here
			if err != nil {
				return fmt.Errorf("error receiving account update: %w", err)
			}
			fmt.Printf("account updated: %+v\n", msg)
		}
	}
}
