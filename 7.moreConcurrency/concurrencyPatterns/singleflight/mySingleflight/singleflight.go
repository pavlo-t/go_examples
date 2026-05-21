package mySingleflight

import (
	"sync"
)

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (any, error)) (any, error, bool) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		c.dups++
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err, true
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()
	g.doCall(c, key, fn)
	return c.val, c.err, c.dups > 0
}

func (g *Group) doCall(c *call, key string, fn func() (any, error)) {
	result, err := fn()
	g.mu.Lock()
	defer g.mu.Unlock()
	c.val = result
	c.err = err
	c.wg.Done()
	if g.m[key] == c {
		delete(g.m, key)
	}
	for _, ch := range c.chans {
		ch <- Result{Val: c.val, Err: c.err, Shared: c.dups > 0}
		close(ch)
	}
}

func (g *Group) Forget(key string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.m, key)
}

type call struct {
	wg    sync.WaitGroup
	val   any
	err   error
	dups  int
	chans []chan<- Result
}

type Result struct {
	Val    any
	Err    error
	Shared bool
}
