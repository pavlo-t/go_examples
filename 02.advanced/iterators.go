package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

type ListIter[T any] struct {
	head, tail *listIterElement[T]
}

type listIterElement[T any] struct {
	next *listIterElement[T]
	val  T
}

func (lst *ListIter[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &listIterElement[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &listIterElement[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *ListIter[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := ListIter[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
