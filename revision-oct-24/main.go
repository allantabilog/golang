package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (u User) String() string {
	return fmt.Sprintf("Hey I'm %v message me at %v", u.FirstName, u.Email)
}

type StrList []string

func (s StrList) String() string {
	return strings.Join(s, ", ")
}

func main() {
	joe := User{ID: 1, FirstName: "Joe", LastName: "Doe", Email: "joe@gmail.com"}
	fmt.Println(joe)

	names := []string{"John", "Jane", "Doe"}
	fmt.Println(StrList(names))
}

func main2() {
	// Marshal
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@gmail.com"}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

}
