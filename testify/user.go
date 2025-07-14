// user.go - User management for demonstrating struct testing and mocking
package main

import (
	"fmt"
	"time"
)

// User represents a user in the system
type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Age      int       `json:"age"`
	IsActive bool      `json:"is_active"`
	Created  time.Time `json:"created"`
}

// NewUser creates a new user
func NewUser(id int, name, email string, age int) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Age:      age,
		IsActive: true,
		Created:  time.Now(),
	}
}

// IsAdult checks if user is 18 or older
func (u *User) IsAdult() bool {
	return u.Age >= 18
}

// Deactivate deactivates the user
func (u *User) Deactivate() {
	u.IsActive = false
}

// Activate activates the user
func (u *User) Activate() {
	u.IsActive = true
}

// GetDisplayName returns formatted display name
func (u *User) GetDisplayName() string {
	if u.Name == "" {
		return "Anonymous User"
	}
	return fmt.Sprintf("%s (%s)", u.Name, u.Email)
}

// Validate validates user data
func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if u.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if u.Age < 0 {
		return fmt.Errorf("age cannot be negative")
	}
	return nil
}
