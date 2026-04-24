package main

import "fmt"

// SlicesIndex takes a slice of any comparable type and an element of that type and
// returns the index of the first occurrence of v in s, or -1 if not present.
// See https://go.dev/blog/deconstructing-type-parameters.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// List - example of a generic type
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push - methods on generic types require type parameters: it's (lst *List[T]), not (lst List)
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}
	// types are inferred:
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
	// types can be specified explicitly:
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
