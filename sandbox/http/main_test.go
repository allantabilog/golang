package main

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		n        int
		expected []int
	}{
		{0, []int{}},
		{1, []int{0}},
		{2, []int{0, 1}},
		{3, []int{0, 1, 1}},
		{4, []int{0, 1, 1, 2}},
		{5, []int{0, 1, 1, 2, 3}},
		{6, []int{0, 1, 1, 2, 3, 5}},
	}

	for _, test := range tests {
		if got := fibonacci(test.n); !reflect.DeepEqual(got, test.expected) {
			t.Errorf("fib(%d) = %d; want %d", test.n, got, test.expected)
		}
	}
}
