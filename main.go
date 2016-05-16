package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}

	tokenizer := New(string(data), true)

	tokenizer.Tokenize()

	fmt.Println(tokenizer.Tokens)
}
