package main

import (
	"sort"
	"testing"

	"pgregory.net/rapid"
)

func TestSortStrings(t *testing.T) {
	rapid.Check(t, func(t *rapid.T){
		s := rapid.SliceOf(rapid.String()).Draw(t, "s")
		sort.Strings(s)
		if !sort.StringsAreSorted(s) {
			t.Fatalf("unsorted after calling sort(): %v", s)
		}
	})
}

func TestArithmetic(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// generate some random integers
		var (
			a = rapid.Int().Draw(t, "a")
			b = rapid.Int().Draw(t, "b")
			c = rapid.Int().Draw(t, "c")
		)

		// test some properties
		if add(a, 0) != a {
			t.Fatal("add() does not have zero as identity")
		}

		if add(a, b) != add(b, a) {
			t.Fatal("add() is not commutative")
		}

		if add(a, add(b, c)) != add(add(a, b), c) {
			t.Fatal("add() is not associative")
		}

	})
}

func add(a, b int) int {
	return (1 + a) - (1 - b)
}

type User struct {
	Name string 
	Age int
}
func TestRapidGenerator(t * testing.T) {
	exampleCount := 10
	
	t.Run("Random Integers", func(t *testing.T){
		for i := range exampleCount {
			randomInt := rapid.Int().Example()
			t.Logf("Random integer #%d: %d", i+1, randomInt)
		}
	})

	t.Run("Random Strings", func(t *testing.T){
		for i := range exampleCount {
			randomString := rapid.String().Example()
			t.Logf("Random string #%d: %s", i+1, randomString)
		}
	})
	
	t.Run("Random structs", func(t *testing.T){
		for i:=range exampleCount {
			randomUser := rapid.Make[*User]()
			t.Logf("Random user: %v", randomUser.Example(i))
		}
	})
}