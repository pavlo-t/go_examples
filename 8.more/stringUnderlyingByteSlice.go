package main

import (
	"fmt"
	"unsafe"
)

func main() {
	suffix := "?"
	// this is necessary because constant strings are stored in read-only memory, mutation would result in a panic
	s1 := "hello" + suffix

	b := unsafe.Slice(unsafe.StringData(s1), len(s1))

	fmt.Println(s1) // hello?
	b[0] = 'H'
	b[len(b)-1] = '!'
	fmt.Println(s1) // Hello!

	fmt.Println("len(s1):", len(s1), "len(b):", len(b), "cap(b):", cap(b))
	b = append(b, '.')
	b[1] = 'o'
	fmt.Println(s1) // Hello!

	s2 := unsafe.String(&b[0], len(b))
	fmt.Println(s2) // Hollo!.
	fmt.Println("len(s2):", len(s2), "len(b):", len(b), "cap(b):", cap(b))

	b = append(b, '.')
	b[3] = 'a'
	fmt.Println(s2) // Holao!.
	fmt.Println("len(s2):", len(s2), "len(b):", len(b), "cap(b):", cap(b))
}
