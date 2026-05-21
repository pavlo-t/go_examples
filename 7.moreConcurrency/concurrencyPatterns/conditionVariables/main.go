package main

import (
	"log"
	"time"

	"go_by_example/7.moreConcurrency/concurrencyPatterns/conditionVariables/myQueueChannels"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// q := myQueueSyncCond.NewQueue[int]()
	// q := myQueueChannels.NewQueueWithoutGetMany[int]()
	q := myQueueChannels.NewQueue[int]()

	i := 1
	pushN := func(n int) {
		time.Sleep(100 * time.Millisecond)
		maxI := n + i
		for ; i <= maxI; i++ {
			time.Sleep(50 * time.Millisecond)
			// log.Printf("Push #%d\n", i)
			q.Push(i)
		}
	}
	go pushN(6)

	go func() { log.Printf("getMany(5): %v\n", q.GetMany(5)) }()
	go func() { log.Printf("getMany(3): %v\n", q.GetMany(3)) }()
	for i := 1; i <= 4; i++ {
		go func() {
			//log.Printf("Get #%d\n", i)
			log.Printf("Get #%d result: %v\n", i, q.Get())
		}()
	}

	for i < 6 {
		time.Sleep(100 * time.Millisecond)
	}
	go pushN(8)

	time.Sleep(500 * time.Millisecond)
}
