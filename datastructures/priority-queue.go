package main

import (
	"container/heap"
	"fmt"
)

// use container/heap package to implement a priority queue
/**
* 1. Define a collection that will implement the heap.Interface
* 2. Use that collection to store the priority queue items
 */

type IntHeap []int 


// make IntHeap implement the heap.Interface
// that means: implement Len, Less, Swap, Push, Pop methods

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func mainPriorityQueue() {
	h := &IntHeap{2, 1, 5, 8, 4, 6, 19, 25, 12, 18, 11}
	heap.Init(h)
	
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("minimum: %d\n", (*h)[0])
		fmt.Printf("Pop: %d \n", heap.Pop(h))
	}
}