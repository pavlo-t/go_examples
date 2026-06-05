package main

// go test -benchmem -bench=. stringUnderlyingByteSlice_test.go
import (
	"testing"
	"unsafe"
)

var (
	suf        = "?"
	testString = `
This is a moderately sized string used to see the difference between standard and unsafe conversions in Go.
` + suf
	testBytes = []byte(testString)
)

func BenchmarkStringToBytesStandardNoMut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// without mutation compiler optimizes away the slice creation
		_ = []byte(testString)
	}
}

func BenchmarkStringToBytesUnsafeNoMut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = unsafe.Slice(unsafe.StringData(testString), len(testString))
	}
}

func BenchmarkStringToBytesStandardMut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// without mutation compiler optimizes away the slice creation
		b := []byte(testString)
		b[0] = 'H'
	}
}

func BenchmarkStringToBytesUnsafeMut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := unsafe.Slice(unsafe.StringData(testString), len(testString))
		b[0] = 'H'
	}
}

func BenchmarkBytesToStringStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(testBytes)
	}
}

func BenchmarkBytesToStringUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = unsafe.String(unsafe.SliceData(testBytes), len(testBytes))
	}
}
