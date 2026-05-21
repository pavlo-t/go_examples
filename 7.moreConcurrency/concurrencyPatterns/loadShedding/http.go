package main

import (
	"fmt"
	"net/http"
	"time"
)

var requestQueue = make(chan struct{}, 3)

func handleRequest(w http.ResponseWriter, _ *http.Request) {
	select {
	case requestQueue <- struct{}{}:
		defer func() { <-requestQueue }()

		time.Sleep(5 * time.Second)
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "Success")

	default:
		w.WriteHeader(http.StatusServiceUnavailable) // HTTP 503
		_, _ = fmt.Fprintln(w, "Server overloaded. Request dropped.")
	}
}

func main() {
	http.HandleFunc("/work", handleRequest)
	fmt.Println("Server running on :8080...")
	_ = http.ListenAndServe(":8080", nil)
}
