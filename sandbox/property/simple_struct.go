package main

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Name struct {
	FirstName string 
	LastName string
}

type UserProfile struct {
	ID int
	Name Name 
	UserName string 
	Password string 
}

func (name *Name) Validate() error {
	if utf8.RuneCountInString(name.FirstName) > 80 {
		return  fmt.Errorf("first name exceeds maximum length of 80 characters")
	}
	if utf8.RuneCountInString(name.LastName) > 80 {
		return  fmt.Errorf("last name exceeds maximum length of 80 characters")
	}
	return nil
}

func (user *UserProfile) Validate() error {
	if err := user.Name.Validate(); err != nil {
		return  err
	}

	if utf8.RuneCountInString(user.UserName) > 80 {
		return  fmt.Errorf("user name exceeds maximum length of 20 characters")
	}

	validUsername := regexp.MustCompile(`^[a-zA-Z0-9\.]+$`).MatchString(user.UserName)
	if !validUsername {
		return fmt.Errorf("username must contain only letters, numbers, dots, and hyphens")
	}
	return nil
}


func (user *UserProfile) String() string {
	return fmt.Sprintf("User %s %s (ID: %d, Username: %s)", user.Name.FirstName, user.Name.LastName, user.ID, user.UserName)
}

