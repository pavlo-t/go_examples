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

var pool sync.Pool

func useObject(jobId int) {
	var o *object
	oFromPool := pool.Get()
	if oFromPool != nil {
		o = oFromPool.(*object)
		log.Printf("job #%d got object from the pool: length: %4d, capacity: %d",
			jobId, len(o.data), cap(o.data))
	} else {
		log.Printf("job #%d failed to get object from the pool, creating new", jobId)
		o = &object{
			data: make([]byte, 0, 1024),
		}
	}
	defer func() {
		o.Reset()
		pool.Put(o)
	}()

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
