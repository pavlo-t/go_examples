package main

import (
	"fmt"
	"sync"
	"time"
)

type Broadcaster[T any] struct {
	mu          sync.RWMutex
	subscribers map[chan T]bool
}

func NewBroadcaster[T any]() *Broadcaster[T] {
	return &Broadcaster[T]{
		subscribers: make(map[chan T]bool),
	}
}

func (b *Broadcaster[T]) Subscribe() chan T {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Using a buffered channel prevents a slow consumer from blocking the broadcaster
	ch := make(chan T, 10)
	b.subscribers[ch] = true
	return ch
}

func (b *Broadcaster[T]) Unsubscribe(ch chan T) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.subscribers[ch]; exists {
		delete(b.subscribers, ch)
		close(ch)
	}
}

func (b *Broadcaster[T]) Broadcast(val T) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for ch := range b.subscribers {
		select {
		case ch <- val: // Send data if buffer has room
		default: // Drop or handle slow consumer to prevent system deadlock
			fmt.Println("Warning: Subscriber buffer full, dropping message")
		}
	}
}

func main() {
	b := NewBroadcaster[string]()

	ch1 := b.Subscribe()
	go func() {
		for msg := range ch1 {
			fmt.Println("Consumer 1 received:", msg)
		}
	}()

	b.Broadcast("After subscribe ch1")
	time.Sleep(50 * time.Millisecond)

	ch2 := b.Subscribe()
	go func() {
		for msg := range ch2 {
			fmt.Println("Consumer 2 received:", msg)
		}
	}()

	b.Broadcast("After subscribe ch2")
	time.Sleep(50 * time.Millisecond)

	b.Unsubscribe(ch1)
	b.Broadcast("After unsubscribe ch1")
	time.Sleep(50 * time.Millisecond)

	b.Unsubscribe(ch2)
	b.Broadcast("After unsubscribe ch2")
	time.Sleep(50 * time.Millisecond)
}
