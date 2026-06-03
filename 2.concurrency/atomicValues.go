package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	type Config struct {
		p1 int
		p2 string
	}

	// suitable for read-heavy, rarely updated data
	var av atomic.Value
	cfg1 := Config{p1: 1, p2: "1"}
	av.Store(cfg1)

	currCfg := av.Load().(Config)
	fmt.Println("config:", currCfg)

	cfg2 := Config{p1: 2, p2: "2"}
	av.Store(cfg2)

	fmt.Println("config:", av.Load())

	cfg3 := Config{p1: 3, p2: "3"}
	old := av.Swap(cfg3)
	fmt.Printf("config after swap: %v, old: %v\n", av.Load(), old)

	cfg3Copy := Config{p1: 3, p2: "3"}
	cfg4 := Config{p1: 4, p2: "4"}
	swapped := av.CompareAndSwap(cfg1, cfg4)
	fmt.Println("swapped:", swapped)
	swapped = av.CompareAndSwap(cfg3Copy, cfg4)
	fmt.Println("swapped:", swapped)

	fmt.Println("config:", av.Load())
}
