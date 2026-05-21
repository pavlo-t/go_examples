package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	numWorkers := 5
	sem := make(chan struct{}, 2)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Go(func() {
			// Acquire a slot (blocks if the buffer is full)
			sem <- struct{}{}

			log.Printf("Worker %d enters slot\n", i)
			time.Sleep(50 * time.Millisecond)
			log.Printf("Worker %d exiting slot\n", i)

			// Release the slot
			<-sem
		})
	}

	wg.Wait()
}
