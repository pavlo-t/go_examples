package myQueueChannels

type Queue[T any] struct {
	s chan state[T]
}

type waiter[T any] struct {
	n int
	c chan []T
}

type state[T any] struct {
	items []T
	wait  []waiter[T]
}

func NewQueue[T any]() *Queue[T] {
	s := make(chan state[T], 1)
	s <- state[T]{}
	return &Queue[T]{s}
}

func (q *Queue[T]) Get() T {
	return q.GetMany(1)[0]
}

func (q *Queue[T]) GetMany(n int) []T {
	s := <-q.s
	if len(s.wait) == 0 && len(s.items) >= n {
		items := s.items[:n:n]
		s.items = s.items[n:]
		q.s <- s
		return items
	}
	c := make(chan []T)
	s.wait = append(s.wait, waiter[T]{n, c})
	q.s <- s
	return <-c
}

func (q *Queue[T]) Push(item T) {
	s := <-q.s
	s.items = append(s.items, item)
	for len(s.wait) > 0 {
		w := s.wait[0]
		if len(s.items) < w.n {
			break
		}
		w.c <- s.items[:w.n:w.n]
		s.items = s.items[w.n:]
		s.wait = s.wait[1:]
	}
	q.s <- s
}
