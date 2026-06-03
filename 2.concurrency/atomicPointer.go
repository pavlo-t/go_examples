package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	type Config struct {
		p1 int
		p2 string
		p3 []int
	}

	var ap atomic.Pointer[Config]
	cfg1 := Config{p1: 1, p2: "1", p3: []int{1}}
	ap.Store(&cfg1)

	currCfg := ap.Load()
	fmt.Println("config:", currCfg)

	cfg2 := Config{p1: 2, p2: "2", p3: []int{2}}
	ap.Store(&cfg2)
	fmt.Println("config:", ap.Load())

	cfg3 := Config{p1: 3, p2: "3", p3: []int{3}}
	old := ap.Swap(&cfg3)
	fmt.Printf("config after swap: %v, old: %v\n", ap.Load(), old)

	cfg3Copy := Config{p1: 3, p2: "3", p3: []int{3}}
	cfg4 := Config{p1: 4, p2: "4", p3: []int{4}}
	swapped := ap.CompareAndSwap(&cfg1, &cfg4)
	fmt.Println("swapped:", swapped)
	swapped = ap.CompareAndSwap(&cfg3Copy, &cfg4)
	fmt.Println("swapped:", swapped)
	swapped = ap.CompareAndSwap(&cfg3, &cfg4)
	fmt.Println("swapped:", swapped)

	fmt.Println("config:", ap.Load())
}
