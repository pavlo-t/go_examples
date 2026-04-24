package main

import (
	"embed"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	// mkdir -p folder
	// echo "hello go" > folder/single_file.txt
	// echo "123" > folder/file1.hash
	// echo "456" > folder/file2.hash

	print(fileString)
	print(string(fileByte))

	content0, _ := folder.ReadFile("folder/single_file.txt")
	print(string(content0))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
