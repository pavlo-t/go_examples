package main

import (
	"log"
	"sync"
	"time"
)

type object struct {
	data []byte
}

func (o *object) Reset() {
	o.data = o.data[:0]
}

var pool = sync.Pool{
	New: func() interface{} {
		return &object{
			data: make([]byte, 0, 1024),
		}
	},
}

func useObject(jobId int) {
	o := pool.Get().(*object)
	defer func() {
		o.Reset()
		pool.Put(o)
	}()

	if cap(o.data) > 1024 {
		log.Printf("job #%d using object from the pool: length: %4d, capacity: %d", jobId, len(o.data), cap(o.data))
	}
	for i := 0; i < 1030; i++ {
		o.data = append(o.data, 'A')
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	for i := 0; i < 10; i++ {
		go useObject(i)
	}

	<-time.After(time.Millisecond * 300)
}
