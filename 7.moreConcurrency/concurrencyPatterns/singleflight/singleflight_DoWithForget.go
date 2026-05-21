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
			key := "key-fetch-data"

			time.Sleep(time.Duration(i) * 40 * time.Millisecond)
			if i%2 == 0 {
				g.Forget(key)
			}
			v, err, shared := g.Do(key, exampleService.FetchData)
			if err != nil {
				log.Fatalf("Error fetching data: %v", err)
			}

			log.Printf("Goroutine %d: result: %v, shared: %v\n", i, v, shared)
		}()
	}

	wg.Wait()
	log.Printf("Function was called %d times\n", exampleService.CallCount.Load())
}
