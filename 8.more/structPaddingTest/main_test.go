package main

// go test -benchmem -bench=. -count=5 8.more/structPaddingTest/main_test.go
import (
	"testing"
)

type Bad struct {
	a bool
	b int
	c bool
}

type Good struct {
	b int
	a bool
	c bool
}

//goland:noinspection GoUnusedGlobalVariable
var (
	SinkBad  []Bad
	SinkGood []Good
)

func BenchmarkBadStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]Bad, 100_000)
		for j := range s {
			s[j].b = j
		}
		SinkBad = s
	}
}

func BenchmarkGoodStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]Good, 100_000)
		for j := range s {
			s[j].b = j
		}
		SinkGood = s
	}
}
