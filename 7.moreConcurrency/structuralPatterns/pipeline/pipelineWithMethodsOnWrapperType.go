package main

import (
	"fmt"
)

type RcvChan[T any] <-chan T

func newRcvChan[T any](els ...T) RcvChan[T] {
	out := make(chan T)
	go func() {
		for _, n := range els {
			out <- n
		}
		close(out)
	}()
	return out
}

type RcvChanInt RcvChan[int]

func newRcvChanInt(nums ...int) RcvChanInt {
	c := newRcvChan(nums...)
	return RcvChanInt(c)
}

func (in RcvChanInt) sq() RcvChanInt {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func (in RcvChanInt) add(i int) RcvChanInt {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + i
		}
		close(out)
	}()
	return out
}

func (in RcvChanInt) mul(i int) RcvChanInt {
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
	c := newRcvChanInt(1, 2, 3)
	out := c.sq()

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println("========================================")

	// We can compose it any way we want:
	for n := range newRcvChanInt(1, 2, 3).sq().sq() {
		fmt.Println(n)
	}
	fmt.Println("========================================")
	for n := range newRcvChanInt(1, 2, 3).sq().mul(2).add(1) {
		fmt.Println(n)
	}
	fmt.Println("========================================")
	for n := range newRcvChanInt(1, 2, 3).add(1).mul(2).sq() {
		fmt.Println(n)
	}
}
