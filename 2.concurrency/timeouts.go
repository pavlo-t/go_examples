package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		log.Println(res)
	case t := <-time.After(1 * time.Second):
		log.Printf("timeout 1, %v\n", t)
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		log.Println(res)
	case t := <-time.After(3 * time.Second):
		log.Printf("timeout 2, %+v\n", t)
	}
}
