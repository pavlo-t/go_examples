package main

import (
	"log"
)

func tee(in <-chan int) (<-chan int, <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range in {
			out1 <- val
			out2 <- val
		}
	}()

	return out1, out2
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	n := 10
	in := make(chan int, n)
	for i := 1; i <= n; i++ {
		in <- i
	}
	close(in)

	out1, out2 := tee(in)

	for i := 0; i < n; i++ {
		log.Printf("Out1: %d, Out2: %d\n", <-out1, <-out2)
	}
}
