package main

import (
	"log"
	"time"
)

type Broker[T any] struct {
	subscribers map[chan T]bool
	publish     chan T
	subscribe   chan chan T
	unsubscribe chan chan T
}

func NewBroker[T any]() *Broker[T] {
	b := &Broker[T]{
		subscribers: make(map[chan T]bool),
		publish:     make(chan T),
		subscribe:   make(chan chan T),
		unsubscribe: make(chan chan T),
	}
	go b.run()
	return b
}

func (b *Broker[T]) run() {
	for {
		select {
		case s := <-b.subscribe:
			b.subscribers[s] = true
		case s := <-b.unsubscribe:
			delete(b.subscribers, s)
			close(s)
		case event := <-b.publish:
			for s := range b.subscribers {
				select {
				case s <- event:
				case <-time.After(100 * time.Millisecond):
					log.Printf("Subscriber too slow, dropping message")
				}
			}
		}
	}
}

func (b *Broker[T]) Broadcast(event T) {
	log.Printf("Publishing event: %T %+q\n", event, event)
	b.publish <- event
}

func (b *Broker[T]) Subscribe() chan T {
	ch := make(chan T, 10)
	b.subscribe <- ch
	return ch
}

func (b *Broker[T]) Unsubscribe(ch chan T) { b.unsubscribe <- ch }

func main() {
	b := NewBroker[string]()

	ch1 := b.Subscribe()
	go func() {
		for msg := range ch1 {
			log.Println("Consumer 1 received:", msg)
		}
	}()

	b.Broadcast("After subscribe ch1")
	time.Sleep(50 * time.Millisecond)

	ch2 := b.Subscribe()
	go func() {
		for msg := range ch2 {
			log.Println("Consumer 2 received:", msg)
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
