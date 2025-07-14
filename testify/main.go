// main.go - Demonstration of the testify learning project
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("=== Testify Learning Project Demo ===")

	// Calculator Demo
	fmt.Println("📱 Calculator Demo:")
	calc := NewCalculator()
	fmt.Printf("5 + 3 = %.2f\n", calc.Add(5, 3))
	fmt.Printf("10 - 4 = %.2f\n", calc.Subtract(10, 4))
	fmt.Printf("Memory: %.2f\n", calc.GetMemory())

	result, err := calc.Divide(10, 2)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}
	fmt.Printf("Is 4 even? %v\n\n", calc.IsEven(4))

	// User Demo
	fmt.Println("👤 User Demo:")
	user := NewUser(1, "John Doe", "john@example.com", 25)
	fmt.Printf("User: %s\n", user.GetDisplayName())
	fmt.Printf("Is adult? %v\n", user.IsAdult())
	fmt.Printf("Is active? %v\n", user.IsActive)

	user.Deactivate()
	fmt.Printf("After deactivation, is active? %v\n\n", user.IsActive)

	// String Utils Demo
	fmt.Println("🔤 String Utils Demo:")
	utils := NewStringUtils()
	fmt.Printf("Reverse 'hello': %s\n", utils.Reverse("hello"))
	fmt.Printf("Is 'racecar' a palindrome? %v\n", utils.IsPalindrome("racecar"))
	fmt.Printf("Word count in 'Hello world test': %d\n", utils.CountWords("Hello world test"))
	fmt.Printf("Capitalize 'hello world': %s\n", utils.Capitalize("hello world"))
	fmt.Printf("Is 'test@example.com' valid email? %v\n", utils.IsValidEmail("test@example.com"))
	fmt.Printf("Initials of 'John Doe Smith': %s\n\n", utils.GetInitials("John Doe Smith"))

	// Service Demo
	fmt.Println("🏢 Service Demo:")
	db := NewInMemoryDatabase()
	email := NewSimpleEmailService()
	userService := NewUserService(db, email)

	newUser, err := userService.CreateUser("Alice Johnson", "alice@example.com", 30)
	if err != nil {
		log.Printf("Error creating user: %v", err)
	} else {
		fmt.Printf("Created user: %s (ID: %d)\n", newUser.GetDisplayName(), newUser.ID)
	}

	// Try to create duplicate user
	_, err = userService.CreateUser("Bob Smith", "alice@example.com", 25)
	if err != nil {
		fmt.Printf("Expected error for duplicate email: %v\n", err)
	}

	fmt.Println("\n🧪 Ready for testing! Run the tests to see testify in action.")
	fmt.Println("Use: go test -v ./...")
}
