package exampleService

import (
	"math/rand"
	"sync/atomic"
	"time"
)

var CallCount atomic.Int32

// FetchData simulates a function that fetches data from a database
func FetchData() (interface{}, error) {
	CallCount.Add(1)
	time.Sleep(100 * time.Millisecond)
	return rand.Intn(100), nil
}
