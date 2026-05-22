package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println("================================================================================")
	s1 := []int{1, 2, 3, 4}
	s2 := s1[:2]   // create new slice with len 2 and capacity as the original (5)
	s3 := s1[:2:2] // create new slice with len 2 and capacity 2
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)
	s3[1] = 5 // mutates s3 and s1 and s2, they all use the same underlying array
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)
	s3 = append(s3, 6) // allocates new array, doesn't mutate s1 and s2
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)
	s2 = append(s2, 7) // mutates both s2 and s1
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)
	s1[1] = 8 // mutates both s1 and s2
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)
}
