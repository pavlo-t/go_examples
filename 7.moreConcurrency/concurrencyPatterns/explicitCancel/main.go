package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				log.Println("Worker stopped!")
				return
			default:
				log.Println("Worker is working...")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	time.Sleep(150 * time.Millisecond)
	log.Println("Sending explicit cancel...")
	close(done)

	time.Sleep(10 * time.Millisecond)
}
