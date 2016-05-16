package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}

	tokenizer := NewTokenizer(string(data), true)

	tokenizer.Tokenize()

	str := ""
	for _, v := range tokenizer.Tokens {
		str += string(v.value)
	}

	fmt.Print(str)
}
