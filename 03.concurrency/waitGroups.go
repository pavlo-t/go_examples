package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// Note that this approach has no straightforward way to propagate errors from workers.
// For more advanced use cases, consider using the https://pkg.go.dev/golang.org/x/sync/errgroup package.
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker2(i)
		})
	}

	wg.Wait()
}
