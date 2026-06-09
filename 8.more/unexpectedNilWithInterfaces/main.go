package main

import "fmt"

type MyError struct{}

func (m *MyError) Error() string {
	return "something went wrong"
}

func getMyError() error { // Returns an iface (error interface)
	var p *MyError = nil
	return p // The interface now has a Type, but Data is nil
}

type MyErrWithAField struct {
	i int
}

func (m *MyErrWithAField) Error() string {
	return fmt.Sprintf("something went wrong")
}

func getMyErrorWithAField() error {
	var p *MyErrWithAField = nil
	return p
}

type MyErrWithAFieldPrinted struct {
	i int
}

func (m *MyErrWithAFieldPrinted) Error() string {
	return fmt.Sprintf("something went wrong; i: %d", m.i)
}

func getMyErrorWithAFieldPrinted() error {
	var p *MyErrWithAFieldPrinted = nil
	return p
}

func main() {
	var err1 *MyError = nil
	if err1 != nil {
		// this won't be printed, err1 is nil
		fmt.Printf("1. err: '%v'\n", err1)
	}

	err2 := getMyError()
	if err2 != nil {
		fmt.Printf("2. err: '%v'\n", err2)
		// 2. err: 'something went wrong'
	}

	err3 := getMyErrorWithAField()
	if err3 != nil {
		fmt.Printf("3. err: '%v'\n", err3)
		// 3. err: 'something went wrong'
	}

	err4 := getMyErrorWithAFieldPrinted()
	if err4 != nil {
		fmt.Printf("4. err: '%v'\n", err4)
		// 4. err: '<nil>'
	}
}
