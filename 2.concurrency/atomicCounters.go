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

	old := ops.Swap(42)
	fmt.Printf("ops after swap: %d, old: %d\n", ops.Load(), old)

	swapped := ops.CompareAndSwap(41, 23)
	fmt.Println("swapped:", swapped)
	swapped = ops.CompareAndSwap(42, 23)
	fmt.Println("swapped:", swapped)
}
