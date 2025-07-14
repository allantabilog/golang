# Testify Learning Project

This project is designed to help you learn the **stretchr/testify** package, one of the most popular testing frameworks for Go. It provides comprehensive examples demonstrating assertions, test suites, mocking, and best practices.

## 🎯 Learning Objectives

By working through this project, you'll learn:
- Basic and advanced assertions using `testify/assert`
- Critical assertions using `testify/require`
- Test suites with setup/teardown using `testify/suite`
- Mocking with `testify/mock`
- Table-driven tests
- Benchmarking
- Testing patterns and best practices

## 📁 Project Structure

```
testify/
├── go.mod                    # Module definition with testify dependency
├── main.go                   # Demo application
├── calculator.go             # Simple calculator for basic testing
├── calculator_test.go        # Demonstrates basic assertions and table-driven tests
├── user.go                   # User model for struct testing
├── user_test.go             # Demonstrates test suites and struct testing
├── service.go               # Service layer with interfaces for mocking
├── service_test.go          # Demonstrates advanced mocking techniques
├── string_utils.go          # String utilities for various test patterns
├── string_utils_test.go     # Demonstrates different assertion types
└── README.md                # This file
```

## 🚀 Getting Started

### 1. Install Dependencies

```bash
cd /Users/atabilog/dev/golang/testify
go mod tidy
```

### 2. Run the Demo Application

```bash
go run .
```

This will demonstrate all the functionality that you'll be testing.

### 3. Run All Tests

```bash
go test -v ./...
```

### 4. Run Specific Test Files

```bash
# Run calculator tests
go test -v -run TestCalculator

# Run user service tests (with mocking)
go test -v -run TestUserService

# Run string utils tests
go test -v -run TestStringUtils
```

### 5. Run Tests with Coverage

```bash
go test -v -cover ./...
```

### 6. Run Benchmarks

```bash
go test -bench=. -v
```

## 📚 Key Testify Features Demonstrated

### 1. Basic Assertions (`assert` package)

**File: `calculator_test.go`**

```go
// Basic equality assertion
assert.Equal(t, 5.0, result, "Addition should return correct result")

// Boolean assertions
assert.True(t, calc.IsEven(2), "2 should be even")
assert.False(t, calc.IsEven(1), "1 should be odd")

// Error handling
assert.NoError(t, err, "Division should not return error for valid input")
assert.Error(t, err, "Division by zero should return error")

// String content assertions
assert.Contains(t, err.Error(), "division by zero")

// Floating point comparisons
assert.InDelta(t, 1.414213, result, 0.000001, "Square root should be approximately correct")
```

### 2. Critical Assertions (`require` package)

**File: `string_utils_test.go`**

```go
// Use require for critical setup - test stops if this fails
require.NotNil(t, utils, "StringUtils instance should not be nil")

// If this fails, the test stops here
result := utils.Reverse("test")
require.NotEmpty(t, result, "Reverse should return non-empty result")

// This only runs if the above require passes
assert.Equal(t, "tset", result)
```

### 3. Test Suites (`suite` package)

**File: `user_test.go`**

```go
type UserTestSuite struct {
    suite.Suite
    user *User
}

// Runs before each test
func (suite *UserTestSuite) SetupTest() {
    suite.user = NewUser(1, "John Doe", "john@example.com", 25)
}

// Runs after each test
func (suite *UserTestSuite) TearDownTest() {
    suite.user = nil
}

func (suite *UserTestSuite) TestUserCreation() {
    assert.Equal(suite.T(), 1, suite.user.ID)
}

// Run the suite
func TestUserTestSuite(t *testing.T) {
    suite.Run(t, new(UserTestSuite))
}
```

### 4. Mocking (`mock` package)

**File: `service_test.go`**

```go
// Create mock
type MockDatabaseService struct {
    mock.Mock
}

func (m *MockDatabaseService) SaveUser(user *User) error {
    args := m.Called(user)
    return args.Error(0)
}

// Set expectations
mockDB.On("SaveUser", mock.AnythingOfType("*main.User")).Return(nil)

// Assert expectations were met
mockDB.AssertExpectations(t)
```

### 5. Table-Driven Tests

**File: `calculator_test.go`**

```go
testCases := []struct {
    name     string
    a, b     float64
    expected float64
}{
    {"positive numbers", 2, 3, 5},
    {"negative numbers", -2, -3, -5},
    {"mixed signs", -2, 3, 1},
}

for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
        result := calc.Add(tc.a, tc.b)
        assert.Equal(t, tc.expected, result)
    })
}
```

## 🎓 Learning Exercises

### Exercise 1: Basic Assertions
1. Look at `calculator_test.go`
2. Add tests for the `Power` method using different assertion types
3. Practice using `assert.InDelta` for floating-point comparisons

### Exercise 2: Test Suites
1. Study `user_test.go`
2. Create a new test suite for the `StringUtils` struct
3. Add setup and teardown methods

### Exercise 3: Mocking
1. Examine `service_test.go`
2. Create a new mock for a payment service interface
3. Practice using custom matchers with `mock.MatchedBy`

### Exercise 4: Table-Driven Tests
1. Convert some of the individual tests in `string_utils_test.go` to table-driven format
2. Practice creating comprehensive test cases

### Exercise 5: Error Testing
1. Add more error scenarios to the service tests
2. Practice testing both expected errors and error handling

## 🔧 Advanced Features to Explore

### Custom Matchers
```go
mockDB.On("SaveUser", mock.MatchedBy(func(user *User) bool {
    return user.Age >= 18 && user.Email != ""
})).Return(nil)
```

### Assertion Functions
```go
func assertValidUser(t *testing.T, user *User) {
    assert.NotNil(t, user)
    assert.NotEmpty(t, user.Name)
    assert.NotEmpty(t, user.Email)
    assert.GreaterOrEqual(t, user.Age, 0)
}
```

### Test Helpers
```go
func createTestUser(t *testing.T, name, email string, age int) *User {
    user := NewUser(1, name, email, age)
    require.NoError(t, user.Validate())
    return user
}
```

## 📊 Running Different Test Types

```bash
# Run only unit tests (exclude integration tests)
go test -v -short ./...

# Run tests with race condition detection
go test -v -race ./...

# Generate coverage report
go test -v -cover -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific test functions
go test -v -run TestCalculatorBasicOperations

# Run tests matching a pattern
go test -v -run "TestUser.*Validation"

# Run benchmarks only
go test -v -run=^$ -bench=.
```

## 🎯 Best Practices Demonstrated

1. **Use descriptive test names** - Each test clearly states what it's testing
2. **Table-driven tests** - Efficient way to test multiple scenarios
3. **Setup and teardown** - Clean test isolation using suites
4. **Mock verification** - Always assert mock expectations
5. **Error testing** - Test both success and failure paths
6. **Edge cases** - Include boundary conditions and edge cases
7. **Benchmarking** - Performance testing for critical functions

## 🔗 Additional Resources

- [Testify Documentation](https://github.com/stretchr/testify)
- [Go Testing Package](https://pkg.go.dev/testing)
- [Table Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Go Testing Best Practices](https://go.dev/doc/tutorial/add-a-test)

## 🎉 Next Steps

After working through this project:
1. Apply these patterns to your own Go projects
2. Explore testify's HTTP testing utilities
3. Learn about property-based testing with [gopter](https://github.com/leanovate/gopter)
4. Study integration testing patterns
5. Explore test doubles and dependency injection patterns

Happy testing! 🧪
