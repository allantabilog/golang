// service.go - Service layer for demonstrating mocking patterns
package main

import (
	"fmt"
	"time"
)

// EmailService interface for sending emails (to be mocked)
type EmailService interface {
	SendEmail(to, subject, body string) error
	SendWelcomeEmail(user *User) error
}

// DatabaseService interface for data persistence (to be mocked)
type DatabaseService interface {
	SaveUser(user *User) error
	GetUser(id int) (*User, error)
	DeleteUser(id int) error
	UserExists(email string) bool
}

// UserService handles business logic for users
type UserService struct {
	db    DatabaseService
	email EmailService
}

// NewUserService creates a new user service
func NewUserService(db DatabaseService, email EmailService) *UserService {
	return &UserService{
		db:    db,
		email: email,
	}
}

// CreateUser creates a new user and sends welcome email
func (s *UserService) CreateUser(name, email string, age int) (*User, error) {
	// Check if user already exists
	if s.db.UserExists(email) {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	// Create user
	user := &User{
		ID:       int(time.Now().Unix()), // Simple ID generation
		Name:     name,
		Email:    email,
		Age:      age,
		IsActive: true,
		Created:  time.Now(),
	}

	// Validate user
	if err := user.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Save to database
	if err := s.db.SaveUser(user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	// Send welcome email
	if err := s.email.SendWelcomeEmail(user); err != nil {
		// Log error but don't fail the user creation
		fmt.Printf("Warning: Failed to send welcome email to %s: %v\n", email, err)
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (*User, error) {
	return s.db.GetUser(id)
}

// DeactivateUser deactivates a user and sends notification
func (s *UserService) DeactivateUser(id int) error {
	user, err := s.db.GetUser(id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	user.Deactivate()

	if err := s.db.SaveUser(user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	// Send deactivation email
	subject := "Account Deactivated"
	body := fmt.Sprintf("Hello %s, your account has been deactivated.", user.Name)
	if err := s.email.SendEmail(user.Email, subject, body); err != nil {
		fmt.Printf("Warning: Failed to send deactivation email: %v\n", err)
	}

	return nil
}

// Real implementations (for demonstration - these would normally be in separate files)

// InMemoryDatabase is a simple in-memory database implementation
type InMemoryDatabase struct {
	users map[int]*User
	emails map[string]bool
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		users:  make(map[int]*User),
		emails: make(map[string]bool),
	}
}

func (db *InMemoryDatabase) SaveUser(user *User) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}
	db.users[user.ID] = user
	db.emails[user.Email] = true
	return nil
}

func (db *InMemoryDatabase) GetUser(id int) (*User, error) {
	user, exists := db.users[id]
	if !exists {
		return nil, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}

func (db *InMemoryDatabase) DeleteUser(id int) error {
	user, exists := db.users[id]
	if !exists {
		return fmt.Errorf("user with id %d not found", id)
	}
	delete(db.emails, user.Email)
	delete(db.users, id)
	return nil
}

func (db *InMemoryDatabase) UserExists(email string) bool {
	return db.emails[email]
}

// SimpleEmailService is a simple email service implementation
type SimpleEmailService struct{}

func NewSimpleEmailService() *SimpleEmailService {
	return &SimpleEmailService{}
}

func (e *SimpleEmailService) SendEmail(to, subject, body string) error {
	fmt.Printf("Sending email to %s: %s\n", to, subject)
	return nil
}

func (e *SimpleEmailService) SendWelcomeEmail(user *User) error {
	subject := "Welcome!"
	body := fmt.Sprintf("Welcome %s! Thanks for joining us.", user.Name)
	return e.SendEmail(user.Email, subject, body)
}
