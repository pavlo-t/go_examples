package main

import (
	"log"
	"time"
)

func process(job int, sem chan struct{}, results chan<- int) {
	defer func() { <-sem }()
	log.Printf("Started job %d\n", job)
	time.Sleep(10 * time.Millisecond)
	results <- job
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	numJobs := 11
	numWorkers := 3

	results := make(chan int)
	sem := make(chan struct{}, numWorkers)

	go func() {
		for res := range results {
			log.Printf("Result: %d\n", res)
		}
	}()

	for j := 1; j <= numJobs; j++ {
		sem <- struct{}{}
		go process(j, sem, results)
	}

	for n := numWorkers; n > 0; n-- {
		sem <- struct{}{}
	}
	close(results)
	close(sem)
}
