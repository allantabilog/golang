package main

import "fmt"

func main() {

	var age int = 47
	var isCool = true
	size := 1.3
	name, email := "allan", "allan.tabilog@gmail.com"

	fmt.Print(name)
	fmt.Println(age, isCool, size, name, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
