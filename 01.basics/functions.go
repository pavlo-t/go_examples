package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func multipleReturnValues() (int, int) {
	return 3, 7
}

func sumVariadic(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeqClosure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func factRecursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * factRecursion(n-1)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// multiple return values
	a, b := multipleReturnValues()
	fmt.Println(a)
	fmt.Println(b)

	_, c := multipleReturnValues()
	fmt.Println(c)

	// variadic arguments
	sumVariadic(1, 2)
	sumVariadic(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sumVariadic(nums...)

	// closures
	nextInt := intSeqClosure()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeqClosure()
	fmt.Println(newInts())

	// recursion
	fmt.Println(factRecursion(7))

	// anonymous recursive functions require explicitly declaring a variable with var
	var fibRecursion func(n int) int
	fibRecursion = func(n int) int {
		if n < 2 {
			return n
		}
		return fibRecursion(n-1) + fibRecursion(n-2)
	}
	fmt.Println(fibRecursion(7))
}
