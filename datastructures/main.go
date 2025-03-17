package main

import (
	"container/heap"
	"fmt"
)

// To use container/heap we first need to
// define a type that implements the heap interface
// that's IntHeap below

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	// next call will make h behave like a heap
	// i.e. it will establish the heap invariant
	heap.Init(h)
	heap.Push(h, -300) // this value becomes the minimum
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("next min: %d\n", heap.Pop(h))
	}
}
