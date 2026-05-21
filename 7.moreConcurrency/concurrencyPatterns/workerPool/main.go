package main

import (
	"log"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		log.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(10 * time.Millisecond)
		results <- job
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	numJobs := 11
	numWorkers := 3

	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Go(func() { worker(i, jobs, results) })
	}

	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		log.Printf("Result: %d\n", res)
	}
}
