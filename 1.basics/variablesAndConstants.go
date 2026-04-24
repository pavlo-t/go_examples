package main

import (
	"fmt"
	"math"
)

const constS string = "constant"

func main() {
	fmt.Println(constS)

	const constN = 500000000

	const constD = 3e20 / constN
	fmt.Println(constD)

	//A numeric constant has no type until it’s given one, such as by an explicit conversion.
	fmt.Println(int64(constD))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call.
	// For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(constN))

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var a1, a2, a3 = 1, 2, "str"
	fmt.Println(a1, a2, a3)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
