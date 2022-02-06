package main

import (
	"fmt"
	"sandbox/maps/sandbox"
)

func main(){
	sandbox.Main()
}


func main0() {
	emails := make(map[string]string)

	emails["Bob"] = "The.Builder@gmail.com"
	emails["Bob2"] = "Another.Builder@gmail.com"
	emails["Bobby"] = "Yet.Another.Builder@gmail.com"

	fmt.Println(emails)
}
