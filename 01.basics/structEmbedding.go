package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	// A container embeds a base. An embedding looks like a field without a name.
	base
	str string
}

// container inherits this if it doesn't have it's own implementation
func (b base) String() string {
	return fmt.Sprintf("Base { num = %v }", b.num)
}

// without this String from base will be used
func (co container) String() string {
	return fmt.Sprintf("Container { base = %v, str = %v }", co.base, co.str)
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Println("Println(co): ", co)
	fmt.Println("Println(co.base): ", co.base)

	// num can be accessed directly on co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// we also can spell out the full path using the embedded type name
	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
