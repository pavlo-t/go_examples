package main

import (
	"fmt"
	"strconv"
	"strings"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func mapC[T, U any](fn func(T) U, in <-chan T) <-chan U {
	out := make(chan U)
	go func() {
		for n := range in {
			out <- fn(n)
		}
		close(out)
	}()
	return out
}

func main() {
	c := gen(1, 2, 3)
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println("========================================")

	// We can compose it any way we want:
	for n := range sq(sq(gen(1, 2, 3))) {
		fmt.Println(n)
	}
	fmt.Println("========================================")
	for n := range sq(double(gen(1, 2, 3))) {
		fmt.Println(n)
	}
	fmt.Println("========================================")
	for n := range double(sq(gen(1, 2, 3))) {
		fmt.Println(n)
	}

	fmt.Println("========================================")
	sq := func(i int) int { return i * i }
	mul := func(i int) func(int) int { return func(j int) int { return i * j } }
	stringify := func(i int) string { return strconv.Itoa(i) }
	repeat := func(i int) func(string) string { return func(s string) string { return strings.Repeat(s, i) } }
	for n := range mapC(repeat(3), mapC(stringify, mapC(sq, mapC(mul(3), gen(1, 2, 3))))) {
		fmt.Println(n)
	}
}
