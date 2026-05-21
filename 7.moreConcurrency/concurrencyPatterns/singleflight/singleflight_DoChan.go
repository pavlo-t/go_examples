package main

import (
	"log"
	"sync"
	"time"

	"go_by_example/7.moreConcurrency/concurrencyPatterns/singleflight/exampleService"
	"golang.org/x/sync/singleflight"
)

//goland:noinspection DuplicatedCode
func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	var wg sync.WaitGroup
	var g singleflight.Group
	// 5 goroutines to fetch the same data
	const numGoroutines = 5
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(i) * 40 * time.Millisecond)
			ch := g.DoChan("key-fetch-data", exampleService.FetchData)
			select {
			case res := <-ch:
				if res.Err != nil {
					log.Fatalf("Error fetching data: %v", res.Err)
				}
				log.Printf("Goroutine %d: result: %v, shared: %v\n", i, res.Val, res.Shared)
			case <-time.After(101 * time.Millisecond):
				log.Fatalf("Timeout waiting for result from goroutine %d", i)
			}
		}()
	}

	wg.Wait()
	log.Printf("Function was called %d times\n", exampleService.CallCount.Load())
}
