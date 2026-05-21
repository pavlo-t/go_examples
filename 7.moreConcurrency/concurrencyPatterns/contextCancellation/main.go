package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("Worker stopped! %v", ctx.Err())
				return
			default:
				log.Println("Worker is working...")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	time.Sleep(150 * time.Millisecond)
	log.Println("Canceling context...")
	cancel()

	time.Sleep(10 * time.Millisecond)
}
