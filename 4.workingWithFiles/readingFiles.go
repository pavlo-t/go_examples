package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Create the file before running:
// echo "hello" > /tmp/dat && echo "go" >> /tmp/dat
//
//goland:noinspection GoUnhandledErrorResult
func main() {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	path := filepath.Join(os.TempDir(), "dat")

	// read whole file into memory
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))

	// open file to perform more operations
	f, err := os.Open(path)
	check(err)

	// read into buffer
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// seek to offset
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// ReadAtLeast will err if fewer than min bytes were read
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// bufio.Reader
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// usually this would be scheduled immediately after Opening with defer
	f.Close()
}
