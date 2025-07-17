// service_test.go - Demonstrates mocking with testify/mock
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// MockDatabaseService is a mock implementation of DatabaseService
type MockDatabaseService struct {
	mock.Mock
}

func (m *MockDatabaseService) SaveUser(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDatabaseService) GetUser(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockDatabaseService) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockDatabaseService) UserExists(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}

// MockEmailService is a mock implementation of EmailService
type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) SendEmail(to, subject, body string) error {
	args := m.Called(to, subject, body)
	return args.Error(0)
}

func (m *MockEmailService) SendWelcomeEmail(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

// UserServiceTestSuite demonstrates mocking with test suites
type UserServiceTestSuite struct {
	suite.Suite
	mockDB    *MockDatabaseService
	mockEmail *MockEmailService
	service   *UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.mockDB = new(MockDatabaseService)
	suite.mockEmail = new(MockEmailService)
	suite.service = NewUserService(suite.mockDB, suite.mockEmail)
}

func (suite *UserServiceTestSuite) TearDownTest() {
	suite.mockDB.AssertExpectations(suite.T())
	suite.mockEmail.AssertExpectations(suite.T())
}

// TestCreateUserSuccess tests successful user creation
func (suite *UserServiceTestSuite) TestCreateUserSuccess() {
	// Setup expectations
	email := "test@example.com"
	suite.mockDB.On("UserExists", email).Return(false)
	suite.mockDB.On("SaveUser", mock.AnythingOfType("*main.User")).Return(nil)
	suite.mockEmail.On("SendWelcomeEmail", mock.AnythingOfType("*main.User")).Return(nil)

	// Execute
	user, err := suite.service.CreateUser("Test User", email, 25)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), user)
	assert.Equal(suite.T(), "Test User", user.Name)
	assert.Equal(suite.T(), email, user.Email)
	assert.Equal(suite.T(), 25, user.Age)
	assert.True(suite.T(), user.IsActive)
}

// TestCreateUserAlreadyExists tests user creation when user already exists
func (suite *UserServiceTestSuite) TestCreateUserAlreadyExists() {
	email := "existing@example.com"
	suite.mockDB.On("UserExists", email).Return(true)

	user, err := suite.service.CreateUser("Test User", email, 25)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Contains(suite.T(), err.Error(), "already exists")
}

// TestCreateUserValidationFailure tests user creation with invalid data
func (suite *UserServiceTestSuite) TestCreateUserValidationFailure() {
	email := "test@example.com"
	suite.mockDB.On("UserExists", email).Return(false)

	// Try to create user with empty name (validation should fail)
	user, err := suite.service.CreateUser("", email, 25)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Contains(suite.T(), err.Error(), "validation failed")
}

// TestCreateUserDatabaseError tests user creation when database save fails
func (suite *UserServiceTestSuite) TestCreateUserDatabaseError() {
	email := "test@example.com"
	dbError := fmt.Errorf("database connection failed")

	suite.mockDB.On("UserExists", email).Return(false)
	suite.mockDB.On("SaveUser", mock.AnythingOfType("*main.User")).Return(dbError)

	user, err := suite.service.CreateUser("Test User", email, 25)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Contains(suite.T(), err.Error(), "failed to save user")
}

// TestCreateUserEmailFailure tests user creation when email sending fails
func (suite *UserServiceTestSuite) TestCreateUserEmailFailure() {
	email := "test@example.com"
	emailError := fmt.Errorf("smtp server down")

	suite.mockDB.On("UserExists", email).Return(false)
	suite.mockDB.On("SaveUser", mock.AnythingOfType("*main.User")).Return(nil)
	suite.mockEmail.On("SendWelcomeEmail", mock.AnythingOfType("*main.User")).Return(emailError)

	// User creation should still succeed even if email fails
	user, err := suite.service.CreateUser("Test User", email, 25)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), user)
}

// TestGetUser tests user retrieval
func (suite *UserServiceTestSuite) TestGetUser() {
	userID := 123
	expectedUser := &User{ID: userID, Name: "Test User", Email: "test@example.com"}

	suite.mockDB.On("GetUser", userID).Return(expectedUser, nil)

	user, err := suite.service.GetUser(userID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUser, user)
}

// TestGetUserNotFound tests user retrieval when user doesn't exist
func (suite *UserServiceTestSuite) TestGetUserNotFound() {
	userID := 999
	dbError := fmt.Errorf("user not found")

	suite.mockDB.On("GetUser", userID).Return((*User)(nil), dbError)

	user, err := suite.service.GetUser(userID)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
}

// TestDeactivateUser tests user deactivation
func (suite *UserServiceTestSuite) TestDeactivateUser() {
	userID := 123
	user := &User{ID: userID, Name: "Test User", Email: "test@example.com", IsActive: true}

	suite.mockDB.On("GetUser", userID).Return(user, nil)
	suite.mockDB.On("SaveUser", mock.MatchedBy(func(u *User) bool {
		return u.ID == userID && !u.IsActive
	})).Return(nil)
	suite.mockEmail.On("SendEmail", user.Email, "Account Deactivated", mock.AnythingOfType("string")).Return(nil)

	err := suite.service.DeactivateUser(userID)

	assert.NoError(suite.T(), err)
	assert.False(suite.T(), user.IsActive)
}

// Run the test suite
func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

// Function-based tests demonstrating different mocking techniques

// TestUserServiceWithCustomMatchers demonstrates custom argument matchers
func TestUserServiceWithCustomMatchers(t *testing.T) {
	mockDB := new(MockDatabaseService)
	mockEmail := new(MockEmailService)
	service := NewUserService(mockDB, mockEmail)

	// Custom matcher for user validation
	mockDB.On("UserExists", "custom@example.com").Return(false)
	mockDB.On("SaveUser", mock.MatchedBy(func(user *User) bool {
		return user.Name == "Custom User" &&
			user.Email == "custom@example.com" &&
			user.Age == 30
	})).Return(nil)

	mockEmail.On("SendWelcomeEmail", mock.MatchedBy(func(user *User) bool {
		return user.Name == "Custom User"
	})).Return(nil)

	user, err := service.CreateUser("Custom User", "custom@example.com", 30)

	require.NoError(t, err)
	assert.NotNil(t, user)
	mockDB.AssertExpectations(t)
	mockEmail.AssertExpectations(t)
}

// TestUserServiceMockCallHistory demonstrates checking mock call history
func TestUserServiceMockCallHistory(t *testing.T) {
	mockDB := new(MockDatabaseService)
	mockEmail := new(MockEmailService)
	service := NewUserService(mockDB, mockEmail)

	// Setup mocks
	mockDB.On("UserExists", mock.Anything).Return(false)
	mockDB.On("SaveUser", mock.Anything).Return(nil)
	mockEmail.On("SendWelcomeEmail", mock.Anything).Return(nil)

	// Create multiple users
	service.CreateUser("User1", "user1@example.com", 25)
	service.CreateUser("User2", "user2@example.com", 30)

	// Assert number of calls
	mockDB.AssertNumberOfCalls(t, "UserExists", 2)
	mockDB.AssertNumberOfCalls(t, "SaveUser", 2)
	mockEmail.AssertNumberOfCalls(t, "SendWelcomeEmail", 2)

	// Assert specific calls were made
	mockDB.AssertCalled(t, "UserExists", "user1@example.com")
	mockDB.AssertCalled(t, "UserExists", "user2@example.com")
}

// TestUserServicePartialMocking demonstrates testing with real implementations
func TestUserServicePartialMocking(t *testing.T) {
	// Use real database but mock email service
	realDB := NewInMemoryDatabase()
	mockEmail := new(MockEmailService)
	service := NewUserService(realDB, mockEmail)

	mockEmail.On("SendWelcomeEmail", mock.AnythingOfType("*main.User")).Return(nil)

	user, err := service.CreateUser("Real DB User", "realdb@example.com", 25)

	require.NoError(t, err)
	assert.NotNil(t, user)

	// Verify user was actually saved in real database
	savedUser, err := realDB.GetUser(user.ID)
	require.NoError(t, err)
	assert.Equal(t, user.Name, savedUser.Name)

	mockEmail.AssertExpectations(t)
}

// BenchmarkUserServiceCreateUser demonstrates benchmarking with mocks
func BenchmarkUserServiceCreateUser(b *testing.B) {
	mockDB := new(MockDatabaseService)
	mockEmail := new(MockEmailService)
	service := NewUserService(mockDB, mockEmail)

	// Setup mocks
	mockDB.On("UserExists", mock.Anything).Return(false)
	mockDB.On("SaveUser", mock.Anything).Return(nil)
	mockEmail.On("SendWelcomeEmail", mock.Anything).Return(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.CreateUser("Benchmark User", fmt.Sprintf("user%d@example.com", i), 25)
	}
}
