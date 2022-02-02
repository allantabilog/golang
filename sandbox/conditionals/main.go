package main

import "fmt"

func main() {
	x := 50
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is greater than or equal %d\n", x, y)
	}

	color := "green"

	if color == "red" {
		fmt.Println("colour is red")
	} else if color == "blue" {
		fmt.Println("colour is blue")
	} else {
		fmt.Println("colour is something else, not blue or red")
	}

	// switch example
	switch color {
	case "red":
		fmt.Println("colour is red")
	case "blue":
		fmt.Println("colour is blue")
	default:
		fmt.Println("colour is not red not blue")
	}

}
