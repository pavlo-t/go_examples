package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker3(id int) error {
	log.Printf("Worker %d starting\n", id)
	if id%2 == 0 {
		return fmt.Errorf("error from worker %d", id)
	}
	time.Sleep(time.Second)
	log.Printf("Worker %d done\n", id)
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	var g errgroup.Group

	for i := 1; i <= 5; i++ {
		g.Go(func() error {
			return worker3(i)
		})
	}

	err := g.Wait()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
