package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

//goland:noinspection GoUnhandledErrorResult
func main() {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	// dump string into a file
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// create and open a file
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	defer f.Close()

	// write a byte slice
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// write a string
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// issue a Sync to flush writes to stable storage
	f.Sync()

	// bufio.Writer
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
