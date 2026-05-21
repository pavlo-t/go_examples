package main

import (
	"log"
	"sync"
	"time"
)

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

func sq(id int, done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		log.Printf("sq %d started", id)
		for n := range in {
			select {
			case out <- n * n:
				log.Printf("sq %d sent %d", id, n*n)
			case <-done:
				log.Printf("sq %d stopped", id)
				return
			}
		}
	}()
	return out
}

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

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// sleep after closing to see messages from stages
	defer func() {
		log.Println("MAIN sleeping in defer")
		time.Sleep(100 * time.Millisecond)
	}()
	done := make(chan struct{})
	defer close(done)

	var nums []int
	for i := 1; i <= 20; i++ {
		nums = append(nums, i)
	}
	in := gen(done, nums...)

	var sqs []<-chan int
	for i := 1; i <= 3; i++ {
		sqs = append(sqs, sq(i, done, in))
	}

	out := merge(done, sqs...)
	log.Println("MAIN received:", <-out)

	time.Sleep(100 * time.Millisecond)
}
