package main

import (
	"fmt"
)

func main() {
	printAllocations := func(s []int) {
		prevCap := cap(s)
		var allocations int
		if prevCap == 0 {
			allocations = 0
		} else {
			allocations = 1
		}
		fmt.Printf("len:% 8d, cap:% 8d, allocations:% 3d\n", len(s), cap(s), allocations)
		for i := 0; i < 1_000_002; i++ {
			s = append(s, i)
			if prevCap != cap(s) {
				prevCap = cap(s)
				allocations += 1
				fmt.Printf("len:% 8d, cap:% 8d, allocations:% 3d\n", len(s), cap(s), allocations)
			}
		}
		fmt.Printf("len:% 8d, cap:% 8d, allocations:% 3d\n", len(s), cap(s), allocations)
	}

	fmt.Println("== Capacity growth: ============================================================")
	var s1 []int
	printAllocations(s1)

	fmt.Println("== Preallocated: ===============================================================")
	s2 := make([]int, 0, 1_000_000)
	printAllocations(s2)
}
