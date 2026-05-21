package myQueueChannels

import "context"

type QueueWithoutGetMany[T any] struct {
	items chan []T  // non-empty slices only
	empty chan bool // holds true if queue is empty
}

func NewQueueWithoutGetMany[T any]() *QueueWithoutGetMany[T] {
	items := make(chan []T, 1)
	empty := make(chan bool, 1)
	empty <- true
	return &QueueWithoutGetMany[T]{items, empty}
}

func (q *QueueWithoutGetMany[T]) Get() T {
	items := <-q.items
	item := items[0]
	if len(items) == 1 {
		q.empty <- true
	} else {
		q.items <- items[1:]
	}
	return item
}

func (q *QueueWithoutGetMany[T]) GetWithCtx(ctx context.Context) (T, error) {
	// just wrap `items := <-q.items` into `select` and add `ctx.Done()` case:
	var items []T
	select {
	case items = <-q.items:
	case <-ctx.Done():
		var item T
		return item, ctx.Err()
	}
	item := items[0]
	if len(items) == 1 {
		q.empty <- true
	} else {
		q.items <- items[1:]
	}
	return item, nil
}

func (q *QueueWithoutGetMany[T]) Push(item T) {
	var items []T
	select {
	case items = <-q.items:
	case <-q.empty:
	}
	items = append(items, item)
	q.items <- items
}
