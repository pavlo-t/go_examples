package main

import "fmt"

func main() {

	// simple for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// with range, equivalent to previous examples
	for i := range 3 {
		fmt.Println("range", i)
	}

	// endless loop with break
	for {
		fmt.Println("loop")
		break
	}

	// loop with continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// range over slice:
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range over map:
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range over string, i is starting byte index of the rune, c is the rune itself
	for i, c := range "go😜😜😜" {
		fmt.Println(i, c)
	}
}
