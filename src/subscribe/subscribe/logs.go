package subscribe

import (
	"context"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func LogsSubscribe(url, program string) error {
	fmt.Println("starting logs subscription...")

	// Connect to Solana WebSocket endpoint
	client, err := ws.Connect(context.Background(), url)
	if err != nil {
		return fmt.Errorf("webSocket connection failed: %w", err)
	}
	defer client.Close()

	// Replace with the program ID you want to track logs for
	programID, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		return fmt.Errorf("invalid program ID: %w", err)
	}

	// Subscribe to logs mentioning the program ID
	sub, err := client.LogsSubscribeMentions(programID, rpc.CommitmentConfirmed)
	if err != nil {
		return fmt.Errorf("log subscription failed: %w", err)
	}
	defer sub.Unsubscribe()

	// set outer timeout (e.g. total listen time)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Listen for log messages
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("subscription timeout: %w", ctx.Err())
		default:
			msg, err := sub.Recv(context.Background()) // no timeout here
			if err != nil {
				return fmt.Errorf("error receiving log update: %w", err)
			}
			fmt.Printf("log received: %+v\n", msg)
		}
	}
}
