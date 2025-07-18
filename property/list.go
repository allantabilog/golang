// list.go - Simple List Data Structure Implementation
package main

import (
	"fmt"
)

// List represents a dynamic array-like data structure
type List struct {
	data     []int // underlying slice to store elements
	size     int   // current number of elements
	capacity int   // current capacity of the underlying array
}

const (
	DefaultInitialCapacity = 10
	GrowthFactor           = 2
	ShrinkThreshold        = 0.25
	ShrinkFactor           = 2
)

// NewList creates a new empty list with default initial capacity
func NewList() *List {
	return &List{
		data:     make([]int, DefaultInitialCapacity),
		size:     0,
		capacity: DefaultInitialCapacity,
	}
}

// NewListWithCapacity creates a new empty list with specified initial capacity
func NewListWithCapacity(initialCapacity int) *List {
	if initialCapacity < 1 {
		initialCapacity = DefaultInitialCapacity
	}
	return &List{
		data:     make([]int, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// Core Operations

// Add appends an element to the end of the list
func (l *List) Add(value int) {
	l.ensureCapacity(l.size + 1)
	l.data[l.size] = value
	l.size++
}

// Insert inserts an element at a specific position
func (l *List) Insert(index int, value int) error {
	if err := l.validateInsertIndex(index); err != nil {
		return err
	}

	l.ensureCapacity(l.size + 1)

	// Shift elements to the right
	for i := l.size; i > index; i-- {
		l.data[i] = l.data[i-1]
	}

	l.data[index] = value
	l.size++
	return nil
}

// Remove removes an element at a specific position
func (l *List) Remove(index int) (int, error) {
	if err := l.validateAccessIndex(index); err != nil {
		return 0, err
	}

	removedValue := l.data[index]

	// Shift elements to the left
	for i := index; i < l.size-1; i++ {
		l.data[i] = l.data[i+1]
	}

	l.size--
	l.shrinkIfNeeded()

	return removedValue, nil
}

// Get retrieves the element at a specific position
func (l *List) Get(index int) (int, error) {
	if err := l.validateAccessIndex(index); err != nil {
		return 0, err
	}
	return l.data[index], nil
}

// Set updates the element at a specific position
func (l *List) Set(index int, value int) error {
	if err := l.validateAccessIndex(index); err != nil {
		return err
	}
	l.data[index] = value
	return nil
}

// Query Operations

// Size returns the number of elements currently in the list
func (l *List) Size() int {
	return l.size
}

// IsEmpty returns true if the list contains no elements
func (l *List) IsEmpty() bool {
	return l.size == 0
}

// Contains returns true if the value exists in the list
func (l *List) Contains(value int) bool {
	for i := 0; i < l.size; i++ {
		if l.data[i] == value {
			return true
		}
	}
	return false
}

// IndexOf returns the first index where the value is found
func (l *List) IndexOf(value int) int {
	for i := 0; i < l.size; i++ {
		if l.data[i] == value {
			return i
		}
	}
	return -1 // Not found
}

// LastIndexOf returns the last index where the value is found
func (l *List) LastIndexOf(value int) int {
	for i := l.size - 1; i >= 0; i-- {
		if l.data[i] == value {
			return i
		}
	}
	return -1 // Not found
}

// Capacity returns the current maximum capacity before resize
func (l *List) Capacity() int {
	return l.capacity
}

// Private helper methods

// ensureCapacity ensures the list can accommodate the required size
func (l *List) ensureCapacity(requiredSize int) {
	if requiredSize > l.capacity {
		l.resize(l.capacity * GrowthFactor)
	}
}

// shrinkIfNeeded reduces capacity if the list is significantly under-utilized
func (l *List) shrinkIfNeeded() {
	if l.capacity > DefaultInitialCapacity &&
		float64(l.size) < float64(l.capacity)*ShrinkThreshold {
		newCapacity := l.capacity / ShrinkFactor
		if newCapacity < DefaultInitialCapacity {
			newCapacity = DefaultInitialCapacity
		}
		l.resize(newCapacity)
	}
}

// resize changes the capacity of the underlying array
func (l *List) resize(newCapacity int) {
	newData := make([]int, newCapacity)
	copy(newData, l.data[:l.size])
	l.data = newData
	l.capacity = newCapacity
}

// validateAccessIndex validates index for access operations (Get, Set, Remove)
func (l *List) validateAccessIndex(index int) error {
	if index < 0 {
		return fmt.Errorf("index %d is negative", index)
	}
	if index >= l.size {
		return fmt.Errorf("index %d out of bounds for size %d", index, l.size)
	}
	return nil
}

// validateInsertIndex validates index for insertion operations
func (l *List) validateInsertIndex(index int) error {
	if index < 0 {
		return fmt.Errorf("index %d is negative", index)
	}
	if index > l.size {
		return fmt.Errorf("index %d out of bounds for insertion in size %d", index, l.size)
	}
	return nil
}

// String returns a string representation of the list
func (l *List) String() string {
	if l.size == 0 {
		return "[]"
	}

	result := "["
	for i := 0; i < l.size; i++ {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%d", l.data[i])
	}
	result += "]"
	return result
}
