package main

import (
	"log"
	"sync"
	"time"
)

type PubSubWithSlices[T any] struct {
	mu     sync.RWMutex
	topics map[string][]chan T
}

func NewPubSubWithSlices[T any]() *PubSubWithSlices[T] {
	return &PubSubWithSlices[T]{
		topics: make(map[string][]chan T),
	}
}

func (ps *PubSubWithSlices[T]) Subscribe(topic string) chan T {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subs, exists := ps.topics[topic]
	if !exists {
		subs = make([]chan T, 0, 10)
	}

	ch := make(chan T, 10)
	ps.topics[topic] = append(subs, ch)
	return ch
}

func (ps *PubSubWithSlices[T]) Unsubscribe(topic string, ch chan T) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subs, topicExists := ps.topics[topic]
	if !topicExists {
		return
	}

	for i, sub := range subs {
		if sub == ch {
			subs[i] = subs[len(subs)-1]
			ps.topics[topic] = subs[:len(subs)-1]
			close(ch)
			break
		}
	}

	if len(subs) == 0 {
		delete(ps.topics, topic)
	}
}

func (ps *PubSubWithSlices[T]) Publish(topic string, val T) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	subs, exists := ps.topics[topic]
	if !exists {
		return
	}

	for _, ch := range subs {
		select {
		case ch <- val: // Successfully delivered message
		default:
			// Non-blocking fallback: drops message if consumer buffer is full
			log.Printf("[Warning] Buffer full for topic %q, dropping message\n", topic)
		}
	}
}

func main() {
	ps := NewPubSubWithSlices[string]()

	sportsSub1 := ps.Subscribe("sports")
	sportsSub2 := ps.Subscribe("sports")

	financeSub := ps.Subscribe("finance")

	// Start consumer routines
	go func() {
		for msg := range sportsSub1 {
			log.Println("[Sports-Sub-1] Received:", msg)
		}
		log.Println("[Sports-Sub-1] Channel Closed Cleanly")
	}()

	go func() {
		for msg := range sportsSub2 {
			log.Println("[Sports-Sub-2] Received:", msg)
		}
		log.Println("[Sports-Sub-2] Channel Closed Cleanly")
	}()

	go func() {
		for msg := range financeSub {
			log.Println("[Finance-Sub]  Received:", msg)
		}
		log.Println("[Finance-Sub]  Channel Closed Cleanly")
	}()

	// 3. Publish initial data streams
	log.Println("--- Round 1: Broadcasting to topics ---")
	ps.Publish("sports", "Goal scored by Team A!")
	ps.Publish("finance", "Stock market reaches new highs.")
	time.Sleep(50 * time.Millisecond)

	// 4. Unsubscribe one of the sports listeners
	log.Println("--- Round 2: Unsubscribing Sports-Sub-1 ---")
	ps.Unsubscribe("sports", sportsSub1)
	time.Sleep(50 * time.Millisecond)

	// 5. Publish again to verify routing changes
	log.Println("--- Round 3: New updates published ---")
	ps.Publish("sports", "Match ended in a draw.")
	ps.Publish("finance", "Interest rates remain unchanged.")
	time.Sleep(50 * time.Millisecond)
}
