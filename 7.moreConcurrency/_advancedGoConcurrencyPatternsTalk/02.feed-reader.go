package main

import (
	"fmt"
	"time"
)

type Item struct {
	Title, Channel, GUID string
}

type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}
type FetcherFunc func() (items []Item, next time.Time, err error)

func (f FetcherFunc) Fetch() (items []Item, next time.Time, err error) {
	return f()
}

func Fetch(domain string) Fetcher {
	return FetcherFunc(func() (items []Item, next time.Time, err error) {
		// TODO generate random number of fake items
		// TODO randomly fail
		// TODO randomly take a long time to answer
		next = time.Now().Add(10 * time.Minute)
		fmt.Printf("Fetching %s: items: %v, next: %v\n", domain, items, next)
		return items, next, err
	})
}

// API of a package provides Fetch, Item, Fetcher, FetchDomain;
// we want a receive-only channel: <-chan Item;
// this is how we transform it into a stream correctly:

type Subscription interface {
	Updates() <-chan Item // stream of Items
	Close() error         // shuts down the stream
}

type sub struct {
	fetcher Fetcher
	updates chan Item
	closing chan chan error
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}
func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}
func (s *sub) loop() {
	type fetchResult struct {
		fetched []Item
		next    time.Time
		err     error
	}
	const maxPending = 10
	var pending []Item // appended by fetch; consumed by send
	var next time.Time // initially January 1, year 0
	var err error
	var seen = make(map[string]bool) // set of item.GUIDs
	var fetchDone chan fetchResult   // if non-nil, Fetch is running
	for {
		var fetchDelay time.Duration // initially 0 (no delay)
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		var startFetch <-chan time.Time
		if fetchDone == nil && len(pending) < maxPending {
			startFetch = time.After(fetchDelay) // enable fetch case
		}

		var first Item
		var updates chan Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates // enable send case
		}

		select {
		case errc := <-s.closing:
			errc <- err
			close(s.updates)
			return
		case <-startFetch:
			fetchDone = make(chan fetchResult, 1)
			go func() {
				fetched, next, err := s.fetcher.Fetch()
				fetchDone <- fetchResult{fetched, next, err}
			}()
		case result := <-fetchDone:
			fetchDone = nil
			if result.err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			next = result.next
			for _, item := range result.fetched {
				if !seen[item.GUID] {
					pending = append(pending, item)
					seen[item.GUID] = true
				}
			}
		case updates <- first:
			pending = pending[1:]
		}
	}
}

// Subscribe converts Fetches to a stream
func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
		closing: make(chan chan error),
	}
	go s.loop()
	return s
}

// Merge merges several streams
func Merge(subs ...Subscription) Subscription {
	// TODO
	if len(subs) == 0 {
		return nil
	}
	return subs[0]
}

func main() {
	// Subscribe to some feeds, and create a merged update stream.
	merged := Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("googleblog.blogspot.com")),
		Subscribe(Fetch("googledevelopers.blogspot.com")),
	)

	// Close the subscriptions after some time.
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	// Print the stream.
	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}

	time.Sleep(50 * time.Millisecond)

	// running with GOTRACEBACK=all will show all goroutine stacks, and leak is visible
	panic("show me the stacks")
}
