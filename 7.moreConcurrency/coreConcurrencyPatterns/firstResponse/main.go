package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// simulateSearch mimics a database query or API call with random latency.
func simulateSearch(ctx context.Context, replicaID int) (string, error) {
	delay := time.Duration(rand.Intn(500)) * time.Millisecond

	select {
	case <-time.After(delay):
		return fmt.Sprintf("Result from replica %d (took %v)", replicaID, delay), nil
	case <-ctx.Done():
		// If another replica won, this context gets canceled to free resources
		return "", ctx.Err()
	}
}

func main() {
	// Create a cancelable context to stop losing goroutines
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx() // Ensure cleanup if not called sooner

	resultChan := make(chan string, 1)

	for i := 1; i <= 3; i++ {
		go func() {
			result, err := simulateSearch(ctx, i)
			if err == nil {
				select {
				case resultChan <- result:
					// Successfully sent the first response!
					// Cancel the context to stop all other replicas immediately
					cancelCtx()
				default:
					// Another goroutine already filled the channel
				}
			}
		}()
	}

	fastestResult := <-resultChan
	fmt.Println("Winner:", fastestResult)

	// Small sleep just to let the cancellation propagate and prove cleanup
	time.Sleep(100 * time.Millisecond)
}
