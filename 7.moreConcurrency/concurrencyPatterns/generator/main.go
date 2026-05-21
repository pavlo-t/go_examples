package main

import "fmt"

func gen(msg string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for {
			c <- msg
		}
	}()
	return c
}

func main() {
	g1 := gen("hello")
	g2 := gen("world")

	for i := 0; i < 10; i++ {
		fmt.Println(<-g1)
		fmt.Println(<-g2)
	}
}
