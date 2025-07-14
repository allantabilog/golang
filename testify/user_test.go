// user_test.go - Demonstrates struct testing and test suites
package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// UserTestSuite demonstrates the testify suite functionality
type UserTestSuite struct {
	suite.Suite
	user *User
}

// SetupTest runs before each test method
func (suite *UserTestSuite) SetupTest() {
	suite.user = NewUser(1, "John Doe", "john@example.com", 25)
}

// TearDownTest runs after each test method
func (suite *UserTestSuite) TearDownTest() {
	suite.user = nil
}

// TestUserCreation tests user creation
func (suite *UserTestSuite) TestUserCreation() {
	assert.Equal(suite.T(), 1, suite.user.ID)
	assert.Equal(suite.T(), "John Doe", suite.user.Name)
	assert.Equal(suite.T(), "john@example.com", suite.user.Email)
	assert.Equal(suite.T(), 25, suite.user.Age)
	assert.True(suite.T(), suite.user.IsActive)
	assert.WithinDuration(suite.T(), time.Now(), suite.user.Created, time.Second)
}

// TestUserIsAdult tests adult checking logic
func (suite *UserTestSuite) TestUserIsAdult() {
	// Current user should be adult
	assert.True(suite.T(), suite.user.IsAdult())

	// Test with minor
	minor := NewUser(2, "Jane Doe", "jane@example.com", 16)
	assert.False(suite.T(), minor.IsAdult())

	// Test edge case - exactly 18
	adult := NewUser(3, "Bob Smith", "bob@example.com", 18)
	assert.True(suite.T(), adult.IsAdult())
}

// TestUserActivation tests activation/deactivation
func (suite *UserTestSuite) TestUserActivation() {
	// Should start as active
	assert.True(suite.T(), suite.user.IsActive)

	// Test deactivation
	suite.user.Deactivate()
	assert.False(suite.T(), suite.user.IsActive)

	// Test reactivation
	suite.user.Activate()
	assert.True(suite.T(), suite.user.IsActive)
}

// TestUserDisplayName tests display name formatting
func (suite *UserTestSuite) TestUserDisplayName() {
	expected := "John Doe (john@example.com)"
	assert.Equal(suite.T(), expected, suite.user.GetDisplayName())

	// Test with empty name
	emptyNameUser := &User{Name: "", Email: "test@example.com"}
	assert.Equal(suite.T(), "Anonymous User", emptyNameUser.GetDisplayName())
}

// TestUserValidation tests user validation
func (suite *UserTestSuite) TestUserValidation() {
	// Valid user should pass validation
	err := suite.user.Validate()
	assert.NoError(suite.T(), err)

	// Test empty name
	invalidUser := &User{Name: "", Email: "test@example.com", Age: 25}
	err = invalidUser.Validate()
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "name cannot be empty")

	// Test empty email
	invalidUser = &User{Name: "John", Email: "", Age: 25}
	err = invalidUser.Validate()
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "email cannot be empty")

	// Test negative age
	invalidUser = &User{Name: "John", Email: "john@example.com", Age: -1}
	err = invalidUser.Validate()
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "age cannot be negative")
}

// Run the test suite
func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

// Standard function-based tests (alternative to suites)

// TestNewUser tests user constructor
func TestNewUser(t *testing.T) {
	id := 123
	name := "Alice Johnson"
	email := "alice@example.com"
	age := 30

	user := NewUser(id, name, email, age)

	assert.NotNil(t, user)
	assert.Equal(t, id, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, age, user.Age)
	assert.True(t, user.IsActive, "New user should be active by default")
	assert.WithinDuration(t, time.Now(), user.Created, time.Second)
}

// TestUserEdgeCases tests edge cases
func TestUserEdgeCases(t *testing.T) {
	t.Run("Zero age user", func(t *testing.T) {
		user := NewUser(1, "Baby", "baby@example.com", 0)
		assert.False(t, user.IsAdult())
		assert.NoError(t, user.Validate())
	})

	t.Run("Very old user", func(t *testing.T) {
		user := NewUser(1, "Elder", "elder@example.com", 120)
		assert.True(t, user.IsAdult())
		assert.NoError(t, user.Validate())
	})

	t.Run("User with special characters in name", func(t *testing.T) {
		user := NewUser(1, "José María", "jose@example.com", 25)
		assert.NoError(t, user.Validate())
		assert.Contains(t, user.GetDisplayName(), "José María")
	})
}

// Table-driven test for IsAdult
func TestUserIsAdultTableDriven(t *testing.T) {
	testCases := []struct {
		name     string
		age      int
		expected bool
	}{
		{"infant", 0, false},
		{"child", 10, false},
		{"teenager", 16, false},
		{"exactly 18", 18, true},
		{"young adult", 25, true},
		{"middle aged", 45, true},
		{"senior", 75, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := NewUser(1, "Test User", "test@example.com", tc.age)
			result := user.IsAdult()
			assert.Equal(t, tc.expected, result,
				"User with age %d should have IsAdult() = %v", tc.age, tc.expected)
		})
	}
}

// TestUserValidationTableDriven demonstrates table-driven validation tests
func TestUserValidationTableDriven(t *testing.T) {
	testCases := []struct {
		name        string
		user        User
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid user",
			user:        User{Name: "John", Email: "john@example.com", Age: 25},
			expectError: false,
		},
		{
			name:        "empty name",
			user:        User{Name: "", Email: "john@example.com", Age: 25},
			expectError: true,
			errorMsg:    "name cannot be empty",
		},
		{
			name:        "empty email",
			user:        User{Name: "John", Email: "", Age: 25},
			expectError: true,
			errorMsg:    "email cannot be empty",
		},
		{
			name:        "negative age",
			user:        User{Name: "John", Email: "john@example.com", Age: -5},
			expectError: true,
			errorMsg:    "age cannot be negative",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.Validate()

			if tc.expectError {
				assert.Error(t, err)
				if tc.errorMsg != "" {
					assert.Contains(t, err.Error(), tc.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
