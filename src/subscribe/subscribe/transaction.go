package subscribe

import (
	"context"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func SignatureSubscribe(url, txHash string) error {
	fmt.Println("starting signature subscription...")

	// Connect to Solana WebSocket endpoint
	client, err := ws.Connect(context.Background(), url)
	if err != nil {
		return fmt.Errorf("websocket connection failed: %w", err)
	}
	defer client.Close()

	// Replace with the signature (transaction hash) you want to track
	sig, err := solana.SignatureFromBase58(txHash)
	if err != nil {
		return fmt.Errorf("invalid signature: %w", err)
	}

	// Subscribe to signature confirmation
	sub, err := client.SignatureSubscribe(sig, rpc.CommitmentConfirmed)
	if err != nil {
		return fmt.Errorf("subscribe error: %w", err)
	}
	defer sub.Unsubscribe()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Wait for confirmation message
	msg, err := sub.Recv(ctx)
	if err != nil {
		return fmt.Errorf("timeout or error receiving signature confirmation: %w", err)

	}
	fmt.Printf("signature confirmed: %+v\n", msg)

	return nil
}
