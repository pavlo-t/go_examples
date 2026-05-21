package main

import (
	"log"
	"sync"
	"time"

	"go_by_example/7.moreConcurrency/concurrencyPatterns/singleflight/exampleService"
	"go_by_example/7.moreConcurrency/concurrencyPatterns/singleflight/mySingleflight"
)

//goland:noinspection DuplicatedCode
func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	var wg sync.WaitGroup
	var g mySingleflight.Group
	// 5 goroutines to fetch the same data
	const numGoroutines = 5
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(i) * 40 * time.Millisecond)
			v, err, shared := g.Do("key-fetch-data", exampleService.FetchData)
			if err != nil {
				log.Fatalf("Error fetching data: %v", err)
			}

			log.Printf("Goroutine %d: result: %v, shared: %v\n", i, v, shared)
		}()
	}

	wg.Wait()
	log.Printf("Function was called %d times\n", exampleService.CallCount.Load())
}
