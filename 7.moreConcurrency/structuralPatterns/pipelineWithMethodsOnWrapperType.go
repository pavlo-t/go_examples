package main

import (
	"fmt"
)

type pipelineChan[T any] <-chan T

func newPipelineChan[T any](els ...T) pipelineChan[T] {
	out := make(chan T)
	go func() {
		for _, n := range els {
			out <- n
		}
		close(out)
	}()
	return out
}

type pipelineChanOfInt pipelineChan[int]

func newPipelineChanOfInt(nums ...int) pipelineChanOfInt {
	c := newPipelineChan(nums...)
	return pipelineChanOfInt(c)
}

func (in pipelineChanOfInt) sq() pipelineChanOfInt {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func (in pipelineChanOfInt) add(i int) pipelineChanOfInt {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + i
		}
		close(out)
	}()
	return out
}

func (in pipelineChanOfInt) mul(i int) pipelineChanOfInt {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * i
		}
		close(out)
	}()
	return out
}

func main() {
	c := newPipelineChanOfInt(1, 2, 3)
	out := c.sq()

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println("========================================")

	// We can compose it any way we want:
	for n := range newPipelineChanOfInt(1, 2, 3).sq().sq() {
		fmt.Println(n)
	}
	fmt.Println("========================================")
	for n := range newPipelineChanOfInt(1, 2, 3).sq().mul(2).add(1) {
		fmt.Println(n)
	}
	fmt.Println("========================================")
	for n := range newPipelineChanOfInt(1, 2, 3).add(1).mul(2).sq() {
		fmt.Println(n)
	}
}
