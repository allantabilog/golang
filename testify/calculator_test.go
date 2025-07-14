// calculator_test.go - Demonstrates basic assertions and table-driven tests
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCalculatorBasicOperations demonstrates basic assertions
func TestCalculatorBasicOperations(t *testing.T) {
	calc := NewCalculator()

	// Basic assertions using assert package
	result := calc.Add(2, 3)
	assert.Equal(t, 5.0, result, "Addition should return correct result")
	assert.Equal(t, 5.0, calc.GetMemory(), "Memory should store last result")

	// Test subtraction
	result = calc.Subtract(10, 4)
	assert.Equal(t, 6.0, result)
	assert.Equal(t, 6.0, calc.GetMemory())

	// Test multiplication
	result = calc.Multiply(3, 4)
	assert.Equal(t, 12.0, result)

	// Test memory operations
	calc.ClearMemory()
	assert.Zero(t, calc.GetMemory(), "Memory should be zero after clearing")
}

// TestCalculatorDivision demonstrates error handling and require vs assert
func TestCalculatorDivision(t *testing.T) {
	calc := NewCalculator()

	// Test normal division
	result, err := calc.Divide(10, 2)
	assert.NoError(t, err, "Division should not return error for valid input")
	assert.Equal(t, 5.0, result)

	// Test division by zero using require (stops test on failure)
	result, err = calc.Divide(10, 0)
	require.Error(t, err, "Division by zero should return error")
	assert.Contains(t, err.Error(), "division by zero")
	assert.Zero(t, result, "Result should be zero when error occurs")
}

// TestCalculatorIsEven demonstrates boolean assertions
func TestCalculatorIsEven(t *testing.T) {
	calc := NewCalculator()

	// Test even numbers
	assert.True(t, calc.IsEven(2), "2 should be even")
	assert.True(t, calc.IsEven(0), "0 should be even")
	assert.True(t, calc.IsEven(-4), "-4 should be even")

	// Test odd numbers
	assert.False(t, calc.IsEven(1), "1 should be odd")
	assert.False(t, calc.IsEven(3), "3 should be odd")
	assert.False(t, calc.IsEven(-3), "-3 should be odd")
}

// TestCalculatorPower demonstrates floating point comparisons
func TestCalculatorPower(t *testing.T) {
	calc := NewCalculator()

	// Test power calculations
	result := calc.Power(2, 3)
	assert.Equal(t, 8.0, result)

	result = calc.Power(5, 2)
	assert.Equal(t, 25.0, result)

	// Test with floating point precision using InDelta
	result = calc.Power(2, 0.5) // Square root of 2
	assert.InDelta(t, 1.414213, result, 0.000001, "Square root should be approximately correct")
}

// Table-driven test structure
type calculatorTestCase struct {
	name     string
	a, b     float64
	expected float64
	hasError bool
}

// TestCalculatorTableDriven demonstrates table-driven tests
func TestCalculatorTableDriven(t *testing.T) {
	calc := NewCalculator()

	additionTests := []calculatorTestCase{
		{"positive numbers", 2, 3, 5, false},
		{"negative numbers", -2, -3, -5, false},
		{"mixed signs", -2, 3, 1, false},
		{"with zero", 5, 0, 5, false},
		{"decimals", 2.5, 1.5, 4.0, false},
	}

	for _, tc := range additionTests {
		t.Run("Addition_"+tc.name, func(t *testing.T) {
			result := calc.Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, result, "Test case: %s", tc.name)
		})
	}

	multiplicationTests := []calculatorTestCase{
		{"positive numbers", 3, 4, 12, false},
		{"with zero", 5, 0, 0, false},
		{"negative numbers", -2, -3, 6, false},
		{"mixed signs", -2, 3, -6, false},
	}

	for _, tc := range multiplicationTests {
		t.Run("Multiplication_"+tc.name, func(t *testing.T) {
			result := calc.Multiply(tc.a, tc.b)
			assert.Equal(t, tc.expected, result, "Test case: %s", tc.name)
		})
	}
}

// TestCalculatorEdgeCases demonstrates testing edge cases
func TestCalculatorEdgeCases(t *testing.T) {
	calc := NewCalculator()

	t.Run("Very large numbers", func(t *testing.T) {
		result := calc.Add(1e10, 1e10)
		assert.Equal(t, 2e10, result)
	})

	t.Run("Very small numbers", func(t *testing.T) {
		result := calc.Add(1e-10, 1e-10)
		assert.InDelta(t, 2e-10, result, 1e-15)
	})

	t.Run("Power of zero", func(t *testing.T) {
		result := calc.Power(0, 5)
		assert.Zero(t, result)
	})
}

// BenchmarkCalculatorAdd demonstrates benchmarking
func BenchmarkCalculatorAdd(b *testing.B) {
	calc := NewCalculator()
	for i := 0; i < b.N; i++ {
		calc.Add(float64(i), float64(i+1))
	}
}

// BenchmarkCalculatorDivide demonstrates benchmarking with error handling
func BenchmarkCalculatorDivide(b *testing.B) {
	calc := NewCalculator()
	for i := 0; i < b.N; i++ {
		calc.Divide(float64(i+1), 2.0) // Avoid division by zero
	}
}
