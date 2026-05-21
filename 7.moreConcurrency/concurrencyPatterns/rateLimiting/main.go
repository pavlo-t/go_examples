package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	requests := make(chan int, 20)
	go func() {
		i := 1
		for ; i <= 5; i++ {
			log.Printf("Sending request %d\n", i)
			requests <- i
		}
		time.Sleep(1500 * time.Millisecond)
		for ; i <= 10; i++ {
			log.Printf("Sending request %d\n", i)
			requests <- i
		}
		close(requests)
	}()

	limiter := time.Tick(100 * time.Millisecond)

	log.Println("Start processing requests")
	for req := range requests {
		<-limiter // Blocks until the 100ms interval passes
		log.Printf("Processed request %d\n", req)
	}

}
