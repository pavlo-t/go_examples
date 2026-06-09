package main

// $ go test -benchmem -bench=. 8.more/allocationBecauseOfInterfaceArg/main_test.go
// BenchmarkConcrete-8     1000000000               0.4516 ns/op          0 B/op          0 allocs/op
// BenchmarkInterface-8    65410300                21.37 ns/op           16 B/op          1 allocs/op
import "testing"

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return u.Name
}

type Guest struct {
	Name string
}

func (g Guest) Greet() string {
	return g.Name
}

var (
	user1 = User{Name: "Alice"}
	user2 = User{Name: "Bob"}
	guest = Guest{Name: "John"}
)

func BenchmarkConcrete(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var u User
		if i%2 == 0 {
			u = user1
		} else {
			u = user2
		}
		result := takeConcrete(u)
		if len(result) == 0 {
			b.Fail()
		}
	}
}

func takeConcrete(u User) string {
	return u.Greet()
}

func BenchmarkInterface(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var u Greeter
		if i%2 == 0 {
			u = user1
		} else {
			u = guest
		}
		result := takeInterface(u)
		if len(result) == 0 {
			b.Fail()
		}
	}
}

func takeInterface(g Greeter) string {
	return g.Greet()
}
