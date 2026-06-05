package main

// $ go test -benchmem -bench=. -count=5 8.more/structPaddingTest/main_test.go
// goos: linux
// goarch: amd64
// cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
// BenchmarkBadStruct-8                6832            220887 ns/op         2400257 B/op          1 allocs/op
// BenchmarkBadStruct-8                3891            264177 ns/op         2400258 B/op          1 allocs/op
// BenchmarkBadStruct-8                4333            232759 ns/op         2400258 B/op          1 allocs/op
// BenchmarkBadStruct-8                5976            220011 ns/op         2400259 B/op          1 allocs/op
// BenchmarkBadStruct-8                5742            232769 ns/op         2400258 B/op          1 allocs/op
// BenchmarkGoodStruct-8               8974            141120 ns/op         1605637 B/op          1 allocs/op
// BenchmarkGoodStruct-8               8574            139439 ns/op         1605638 B/op          1 allocs/op
// BenchmarkGoodStruct-8               7201            146399 ns/op         1605638 B/op          1 allocs/op
// BenchmarkGoodStruct-8               8566            132990 ns/op         1605639 B/op          1 allocs/op
// BenchmarkGoodStruct-8               8626            135354 ns/op         1605639 B/op          1 allocs/op
// BenchmarkBadStructNoAlloc-8        16227             73360 ns/op             147 B/op          0 allocs/op
// BenchmarkBadStructNoAlloc-8        16396             72669 ns/op             146 B/op          0 allocs/op
// BenchmarkBadStructNoAlloc-8        16567             72363 ns/op             144 B/op          0 allocs/op
// BenchmarkBadStructNoAlloc-8        16537             72414 ns/op             145 B/op          0 allocs/op
// BenchmarkBadStructNoAlloc-8        16376             72885 ns/op             146 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       19412             61434 ns/op              82 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       19485             62043 ns/op              82 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       19420             60199 ns/op              82 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       19279             65251 ns/op              83 B/op          0 allocs/op
// BenchmarkGoodStructNoAlloc-8       18819             67358 ns/op              85 B/op          0 allocs/op
// PASS
// ok      command-line-arguments  15.834s
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

func BenchmarkBadStructNoAlloc(b *testing.B) {
	s := make([]Bad, 100_000)
	for i := 0; i < b.N; i++ {
		for j := range s {
			s[j].a = !s[j].a
			s[j].b = j
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
			s[j].b = j
			s[j].c = !s[j].c
		}
		SinkGood = s
	}
}
