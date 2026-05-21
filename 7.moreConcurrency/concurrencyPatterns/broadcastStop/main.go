package main

import (
	"log"
	"sync"
	"time"
)

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
				log.Println("merge sent:", n)
			case <-done:
				log.Println("merge stopped")
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			select {
			case out <- n:
				log.Println("gen sent:", n)
			case <-done:
				log.Println("gen stopped")
				return
			}
		}
		close(out)
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
				log.Println("sq sent:", n*n)
			case <-done:
				log.Println("sq stopped")
				return
			}
		}
	}()
	return out
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// sleep after closing to see messages from stages
	defer func() {
		log.Println("main sleeping in defer")
		time.Sleep(200 * time.Millisecond)
	}()
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 1, 2, 3, 4, 5, 6, 7)

	// Distribute the sq work across multiple goroutines that all read from in.
	c1 := sq(done, in)
	c2 := sq(done, in)
	c3 := sq(done, in)

	// Consume the first value from output.
	out := merge(done, c1, c2, c3)
	log.Println("main received:", <-out) // 1 or 4 or 9
}
