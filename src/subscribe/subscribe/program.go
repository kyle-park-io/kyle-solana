package subscribe

import (
	"context"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func ProgramSubscribe(url, program string) error {
	fmt.Println("starting program subscription...")

	// Connect to Solana WebSocket endpoint
	client, err := ws.Connect(context.Background(), url)
	if err != nil {
		return fmt.Errorf("webSocket connection failed: %w", err)
	}
	defer client.Close()

	// Replace with the program ID whose related accounts you want to monitor
	programID, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		return fmt.Errorf("invalid program ID: %w", err)
	}

	// Subscribe to account changes for the given program
	sub, err := client.ProgramSubscribe(programID, rpc.CommitmentConfirmed)
	if err != nil {
		return fmt.Errorf("program subscribe error: %w", err)
	}
	defer sub.Unsubscribe()

	// set outer timeout (e.g. total listen time)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Listen for updates to any accounts owned by the program
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("subscription timeout: %w", ctx.Err())
		default:
			msg, err := sub.Recv(context.Background()) // no timeout here
			if err != nil {
				return fmt.Errorf("error receiving program update: %w", err)
			}
			fmt.Printf("program updated: %+v\n", msg)
		}
	}
}
