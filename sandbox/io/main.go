package main

import (
	"fmt"
	"os"
)

func main() {
	var filename = "/Users/allantabilog/dev/golang/sandbox/io/main.go"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}