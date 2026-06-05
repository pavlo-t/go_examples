package main

// $ go test -benchmem -bench=. -count=5 8.more/structPaddingTest/main_test.go
// goos: linux
// goarch: amd64
// cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
// BenchmarkBadStruct-8                6433            215287 ns/op         2400258 B/op          1 allocs/op
// BenchmarkBadStruct-8                6214            181616 ns/op         2400257 B/op          1 allocs/op
// BenchmarkBadStruct-8                5625            194506 ns/op         2400257 B/op          1 allocs/op
// BenchmarkBadStruct-8                6594            191925 ns/op         2400257 B/op          1 allocs/op
// BenchmarkBadStruct-8                6630            185213 ns/op         2400257 B/op          1 allocs/op
// BenchmarkGoodStruct-8               8433            123712 ns/op         1605635 B/op          1 allocs/op
// BenchmarkGoodStruct-8               9664            122348 ns/op         1605635 B/op          1 allocs/op
// BenchmarkGoodStruct-8               9788            128095 ns/op         1605635 B/op          1 allocs/op
// BenchmarkGoodStruct-8               9193            125675 ns/op         1605635 B/op          1 allocs/op
// BenchmarkGoodStruct-8               9768            124760 ns/op         1605636 B/op          1 allocs/op
// BenchmarkBadStructNoAlloc-8        16142             80697 ns/op             148 B/op          0 allocs/op
// BenchmarkBadStructNoAlloc-8        15387             75834 ns/op             155 B/op          0 allocs/op
// BenchmarkBadStructNoAlloc-8        13615             75204 ns/op             176 B/op          0 allocs/op
// BenchmarkBadStructNoAlloc-8        15715             73439 ns/op             152 B/op          0 allocs/op
// BenchmarkBadStructNoAlloc-8        16353             74815 ns/op             146 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       19194             61902 ns/op              83 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       19387             61776 ns/op              82 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       19545             61578 ns/op              82 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       18829             70859 ns/op              85 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       19213             63179 ns/op              83 B/op          0 allocs/op
// PASS
// ok      command-line-arguments  35.327s
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
			s[j].b = i
		}
		SinkBad = s
	}
}

func BenchmarkGoodStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]Good, 100_000)
		for j := range s {
			s[j].b = i
		}
		SinkGood = s
	}
}

func BenchmarkBadStructNoAlloc(b *testing.B) {
	s := make([]Bad, 100_000)
	for i := 0; i < b.N; i++ {
		for j := range s {
			s[j].a = !s[j].a
			s[j].b = i
			s[j].c = !s[j].c
		}
		SinkBad = s
	}
}

func BenchmarkGoodStructNoAlloc(b *testing.B) {
	s := make([]Good, 100_000)
	for i := 0; i < b.N; i++ {
		for j := range s {
			s[j].a = !s[j].a
			s[j].b = i
			s[j].c = !s[j].c
		}
		SinkGood = s
	}
}
