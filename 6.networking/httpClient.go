package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// ----------------------------------------------------------------------------------------------------

	req, err := http.NewRequest("GET", "https://gobyexample.com", nil)
	if err != nil {
		panic(err)
	}
	//req.Header.Add("Hello", "World")
	fmt.Printf("req: %+v\n", req)
	resp2, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()
	fmt.Printf("resp: %+v\n", resp2)

	scanner2 := bufio.NewScanner(resp2.Body)
	for i := 0; scanner2.Scan() && i < 5; i++ {
		fmt.Println(scanner2.Text())
	}

	if err := scanner2.Err(); err != nil {
		panic(err)
	}
}
