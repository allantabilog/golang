package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}	

// ByAge implements sort.Interface for []Person based on Age
type ByAge []Person 

func (a ByAge) Len() int 			{ return len(a) }
func (a ByAge) Swap(i, j int)		{ a[i], a[j] = a[j], a[i]}
func (a ByAge) Less(i, j int) bool	{ return a[i].Age < a[j].Age} 

type User struct {
	ID int
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("%d: %s", u.ID, u.Name)
}

func main() {
	people := []Person {
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)

	users := []User {
		{ID: 1, Name: "Bob"},
		{ID: 2, Name: "Katy"},
	}
	var joe User
	joe = User{ID: 3, Name: "Joe"}
	fmt.Println(users)
	fmt.Println(joe)

	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	fmt.Println(people)
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)
}


