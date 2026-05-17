package main

import (
	"log"
	"sync"
	"time"
)

type PubSub[T any] struct {
	mu     sync.RWMutex
	topics map[string]map[chan T]bool
}

func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{
		topics: make(map[string]map[chan T]bool),
	}
}

func (ps *PubSub[T]) Subscribe(topic string) chan T {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, exists := ps.topics[topic]; !exists {
		ps.topics[topic] = make(map[chan T]bool)
	}

	// Buffer of 10 prevents slow consumers from instantly blocking publishers
	ch := make(chan T, 10)
	ps.topics[topic][ch] = true
	return ch
}

func (ps *PubSub[T]) Unsubscribe(topic string, ch chan T) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subs, topicExists := ps.topics[topic]
	if !topicExists {
		return
	}

	if _, chExists := subs[ch]; chExists {
		delete(subs, ch)
		close(ch)
	}

	if len(subs) == 0 {
		delete(ps.topics, topic)
	}
}

func (ps *PubSub[T]) Publish(topic string, val T) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	subs, exists := ps.topics[topic]
	if !exists {
		return
	}

	for ch := range subs {
		select {
		case ch <- val: // Successfully delivered message
		default:
			// Non-blocking fallback: drops message if consumer buffer is full
			log.Printf("[Warning] Buffer full for topic %q, dropping message\n", topic)
		}
	}
}

func main() {
	ps := NewPubSub[string]()

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
