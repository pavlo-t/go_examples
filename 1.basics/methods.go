package main

import "fmt"

type methodsRect struct {
	width, height int
}

// use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the struct
func (r *methodsRect) area() int {
	return r.width * r.height
}

func (r methodsRect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := methodsRect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// go automatically handles conversion between values and pointers for method calls
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
