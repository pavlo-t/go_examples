package main

import (
	"log"
	"time"

	"go_by_example/7.moreConcurrency/concurrencyPatterns/conditionVariables/myQueue"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	q := myQueue.NewQueue[int]()
	for i := 1; i <= 3; i++ {
		go func() {
			log.Printf("Get #%d\n", i)
			v := q.Get()
			log.Printf("Get #%d result: %v\n", i, v)
		}()
	}
	go func() {
		time.Sleep(100 * time.Millisecond)
		for i := 1; i <= 10; i++ {
			q.Push(i)
		}
	}()

	time.Sleep(50 * time.Millisecond)

	log.Printf("Main get: %v\n", q.Get())
	log.Printf("Main get: %v\n", q.Get())

	time.Sleep(100 * time.Millisecond)
}
