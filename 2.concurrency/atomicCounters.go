package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64
	i := 1
	var wg sync.WaitGroup

	for range 50 {
		wg.Go(func() {
			for range 1000 {
				ops.Add(1)
				i++
			}
		})
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
	fmt.Println("i:", i)
}
