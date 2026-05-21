package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	ongoingRequests := make(chan struct{}, 3)

	handleReq := func(id int) {
		select {
		case ongoingRequests <- struct{}{}:
			defer func() { <-ongoingRequests }()
			time.Sleep(100 * time.Millisecond)
			log.Printf("Success %d", id)
		default:
			log.Printf("Work queue full, skipping %d", id)
		}
	}

	i := 1
	for ; i <= 5; i++ {
		go handleReq(i)
	}
	time.Sleep(101 * time.Millisecond)
	for ; i <= 10; i++ {
		go handleReq(i)
	}
	time.Sleep(101 * time.Millisecond)
}
