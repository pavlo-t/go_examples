package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

//goland:noinspection GoUnhandledErrorResult
func main() {
	type response1 struct {
		Page   int
		Fruits []string
	}

	type response2 struct {
		Page   int      `json:"page"`
		Fruits []string `json:"fruits"`
	}

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)

	// ----------------------------------------------------------------------------------------------------
	// from https://go.dev/blog/json
	fmt.Println(strings.Repeat("-", 100))
	type Command struct {
		Id    int
		Value string
	}
	type Message struct {
		Value string
	}
	type IncomingMessage struct {
		Cmd Command
		Msg Message
	}
	type IncomingMessagePointers struct {
		Cmd *Command
		Msg *Message
	}

	// version with pointers doesn't allocate missing fields
	incomingCmdEmpty := []byte("")
	var im1 IncomingMessage
	var imp1 IncomingMessagePointers
	json.Unmarshal(incomingCmdEmpty, &im1)
	json.Unmarshal(incomingCmdEmpty, &imp1)
	fmt.Println(im1)
	fmt.Println(imp1)

	incomingCmd := []byte(`{"cmd": {"id": 42, "value": "Don't panic"}}`)
	var im2 IncomingMessage
	var imp2 IncomingMessagePointers
	json.Unmarshal(incomingCmd, &im2)
	json.Unmarshal(incomingCmd, &imp2)
	fmt.Println(im2)
	fmt.Println(imp2)

	incomingMsg := []byte(`{"msg": {"value": "Don't panic"}}`)
	var im3 IncomingMessage
	var imp3 IncomingMessagePointers
	json.Unmarshal(incomingMsg, &im3)
	json.Unmarshal(incomingMsg, &imp3)
	fmt.Println(im3)
	fmt.Println(imp3)

	// ----------------------------------------------------------------------------------------------------
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println(`write json with "Name" field in it, e.g. {"Name":"John"}`)
	decStdIn := json.NewDecoder(os.Stdin)
	encStdOut := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := decStdIn.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := encStdOut.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
