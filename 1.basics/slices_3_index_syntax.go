package main

import (
	"fmt"
)

func main() {
	printSs := func(ss ...[]int) {
		for i, s := range ss {
			fmt.Printf("s%d: %v cap: %d \n", i+1, s, cap(s))
		}
		fmt.Println()
	}

	s1 := []int{1, 2, 3, 4}
	s2 := s1[:2]    // create new slice with len 2 and capacity as the original (4)
	s3 := s1[:2:2]  // create new slice with len 2 and capacity 2
	s4 := s1[1:3:3] // create new slice with len 2 and capacity 2 starting at index 1
	println("Init:")
	printSs(s1, s2, s3, s4)

	// mutates s3 and s1 and s2, they all use the same underlying array
	s3[1] = 5
	println("> s3[1] = 5")
	printSs(s1, s2, s3, s4)

	// allocates new array, doesn't mutate s1 and s2
	s3 = append(s3, 6)
	println("> s3 = append(s3, 6)")
	printSs(s1, s2, s3, s4)

	// mutates both s2 and s1
	s2 = append(s2, 7)
	println("> s2 = append(s2, 7)")
	printSs(s1, s2, s3, s4)

	// mutates both s1 and s2
	s1[1] = 8
	println("> s1[1] = 8")
	printSs(s1, s2, s3, s4)
}
