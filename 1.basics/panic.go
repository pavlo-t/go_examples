package main

import (
	"os"
	"path/filepath"
)

//goland:noinspection GoUnreachableCode
func main() {
	panic("a problem")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
}
