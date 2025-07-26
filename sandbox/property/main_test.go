package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return  a - b
}

// Property: addition is commutative
func TestAddCommutative(t *testing.T) {
	f := func(a, b int) bool {
		return Add(a, b) == Add(b, a)
	}
	 
	if err := quick.Check(f, nil); err != nil {
		t.Errorf("Property failed: %v", err)
	}
}

func TestSubtractPropertyFour(t *testing.T) {
	property := func(a, b int) bool {
		if b > 0 {
			return true
		}
		return Subtract(a, b) > a
	}

	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}


func mainTest(){
	fmt.Println("Running tests...")
}