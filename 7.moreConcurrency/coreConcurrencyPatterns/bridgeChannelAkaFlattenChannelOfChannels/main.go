package main

import (
	"fmt"
)

// Bridge flattens a channel of channels into a single stream of values
func Bridge[T any](done <-chan struct{}, cc <-chan (<-chan T)) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for {
			var c <-chan T
			select {
			case maybeC, ok := <-cc:
				if !ok {
					return
				}
				c = maybeC
			case <-done:
				return
			}

			for val := range c {
				select {
				case out <- val:
				case <-done:
					return
				}
			}
		}
	}()

	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	cc := make(chan (<-chan int))

	go func() {
		defer close(cc)
		makeChan := func(values ...int) <-chan int {
			c := make(chan int)
			go func() {
				defer close(c)
				for _, v := range values {
					c <- v
				}
			}()
			return c
		}

		cc <- makeChan(0, 1, 2)
		cc <- makeChan(3)
		cc <- makeChan(4, 5, 6)
		cc <- makeChan(7, 8, 9)
	}()

	flatStream := Bridge(done, cc)

	for val := range flatStream {
		fmt.Printf("%v ", val)
	}
}
