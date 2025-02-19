package main

import "fmt"

func Greet(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Error: name is empty")
	}
	return "Hello " + name, nil
}

func main() {
	message, err := Greet("John")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(message)

	message, err = Greet("")
	if err != nil {
		fmt.Println(err)
		return
	}
}
