package main

import (
	"fmt"
	"sync"
)

type Container1 struct {
	mu       sync.RWMutex
	counters map[string]int
}

func (c *Container1) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container1{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	wg.Go(func() {
		doIncrement("a", 10000)
	})
	wg.Go(func() {
		doIncrement("a", 10000)
	})
	wg.Go(func() {
		doIncrement("b", 10000)
	})

	wg.Wait()

	c.mu.RLock()
	fmt.Println("main RLock")

	fmt.Println(c.counters)

	for i := range 10 {
		wg.Go(func() {
			c.mu.RLock()
			fmt.Println(i, "go RLock")
			fmt.Println(i, c.counters)
			fmt.Println(i, "go RUnlock")
			c.mu.RUnlock()
		})
	}
	wg.Wait()

	fmt.Println("main RUnlock")
	c.mu.RUnlock()
}
