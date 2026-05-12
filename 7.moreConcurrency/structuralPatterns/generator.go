package main

import "fmt"

func generator(msg string) <-chan string {
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
	g1 := generator("hello")
	g2 := generator("world")

	for i := 0; i < 10; i++ {
		fmt.Println(<-g1)
		fmt.Println(<-g2)
	}
}
