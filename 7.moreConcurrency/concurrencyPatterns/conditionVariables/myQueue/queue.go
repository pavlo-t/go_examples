package myQueue

import "sync"

type Queue[T any] struct {
	mu        sync.Mutex
	items     []T
	itemAdded sync.Cond
}

func NewQueue[T any]() *Queue[T] {
	q := new(Queue[T])
	q.itemAdded.L = &q.mu
	return q
}

func (q *Queue[T]) Get() T {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.items) == 0 {
		// Wait atomically unlocks the mutex and suspends the goroutine
		q.itemAdded.Wait()
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) Push(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
	// Signal relocks the waited mutex and proceeds with execution
	q.itemAdded.Signal()
}
