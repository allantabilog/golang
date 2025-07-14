// calculator.go - Simple calculator for basic testing
package main

import (
	"errors"
	"math"
)

// Calculator represents a simple calculator
type Calculator struct {
	memory float64
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{memory: 0}
}

// Add performs addition
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.memory = result
	return result
}

// Subtract performs subtraction
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.memory = result
	return result
}

// Multiply performs multiplication
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.memory = result
	return result
}

// Divide performs division with error handling
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	result := a / b
	c.memory = result
	return result, nil
}

// Power calculates a to the power of b
func (c *Calculator) Power(a, b float64) float64 {
	result := math.Pow(a, b)
	c.memory = result
	return result
}

// GetMemory returns the last calculated value
func (c *Calculator) GetMemory() float64 {
	return c.memory
}

// ClearMemory resets the memory to zero
func (c *Calculator) ClearMemory() {
	c.memory = 0
}

// IsEven checks if a number is even
func (c *Calculator) IsEven(n int) bool {
	return n%2 == 0
}
