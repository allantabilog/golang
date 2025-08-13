package main

import (
	"fmt"
)

func fibonacci() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		a, b := 0, 1
		for {
			out <- a 
			a, b = b, a + b 
		}
	}()
	return out
}

func fibonacciMain() {
	fibs := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", <-fibs)
	}
}