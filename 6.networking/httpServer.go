package main

import (
	"fmt"
	"net/http"
)

// curl localhost:8090/hello
func httpServerHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// curl localhost:8090/headers
func httpServerHeaders(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// curl localhost:8090/users/Username/hello
	httpServerUsersHello := func(w http.ResponseWriter, req *http.Request) {
		user := req.PathValue("user")
		fmt.Fprintf(w, "Hello, %s!\n\n", user)
	}

	http.HandleFunc("/hello", httpServerHello)
	http.HandleFunc("/headers", httpServerHeaders)
	http.HandleFunc("/users/{user}/hello", httpServerUsersHello)

	http.ListenAndServe(":8090", nil)
}
