package main

import "fmt"

func zeroval(ival int) {
	// ival is a copy of ival distinct from the one in the calling function
	ival = 0
}

func zeroptr(iptr *int) {
	// assigning a value to a dereferenced pointer changes the value at the referenced address
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i gets a pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
