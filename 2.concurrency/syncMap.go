package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("user_101", "Alice")
	m.Store("user_102", "Bob")

	if val, ok := m.Load("user_101"); ok {
		username := val.(string)
		fmt.Println("Loaded user:", username)
	}

	actual, loaded := m.LoadOrStore("user_101", "Charlie")
	fmt.Printf("actual: %v, loaded: %v\n", actual, loaded)

	actual, loaded = m.LoadOrStore("user_103", "Charlie")
	fmt.Printf("actual: %v, loaded: %v\n", actual, loaded)

	m.Range(func(key, value any) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})

	deleted, loaded := m.LoadAndDelete("user_101")
	fmt.Printf("deleted: %v, loaded: %v\n", deleted, loaded)
}
