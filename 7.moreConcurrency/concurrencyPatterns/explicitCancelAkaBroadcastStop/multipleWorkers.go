package main

import (
	"log"
	"sync"
	"time"
)

func worker(id int, done chan struct{}) {
	for {
		select {
		case <-done:
			log.Printf("Worker %d stopped\n", id)
			return
		default:
			log.Printf("Worker %d working\n", id)
			time.Sleep(51 * time.Millisecond)
		}
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	var wg sync.WaitGroup
	done := make(chan struct{})

	for i := 1; i <= 5; i++ {
		wg.Go(func() { worker(i, done) })
	}

	time.Sleep(50 * time.Millisecond)
	log.Println("Broadcasting stop signal...")
	close(done)

	wg.Wait()
}
