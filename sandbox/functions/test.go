package main

import "fmt"

func greeting(name string) string {
	return "Hello.there" + name
}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Brad.Pitt"))
	fmt.Println(add(3, 4))
}
