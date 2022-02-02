package main

import (
	"fmt"
	"strconv"
)

// Define a person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// methods
// a value receiver:
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " age: " + strconv.Itoa(p.age)
}

// a pointer receiver
func (p *Person) incrementAge() {
	p.age++
}

func (p *Person) updateLastName(newLastName string) {
	p.lastName = newLastName
}

func main() {
	// Init person struct
	person1 := Person{firstName: "Bob", lastName: "The.Builder", city: "Melbz", age: 30, gender: "M"}
	fmt.Println(person1)

	person2 := Person{"Bobb", "The.Bbuilder", "Melbourne", "M", 30}
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	person1.incrementAge()
	person1.incrementAge()
	person1.incrementAge()
	person1.incrementAge()
	fmt.Println(person1.greet())

	person1.updateLastName("New.Builder")
	fmt.Println(person1.greet())

}
