package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
	Age  int
}



func main() {
	switchExample3()
}

func getStructValue() User {
	return User{
		Name: "John",
		Age:  30,
	}
}

func switchExample3() {
	t := time.Now()
	fmt.Printf("The time is %v\n", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
func switchExample2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Printf("today is %v\n", today)
	fmt.Printf("today + 1 is %v\n", today + 1)
	fmt.Printf("today + 2 is%v\n", today + 2)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}