package main

import "fmt"

type OffspringProb struct {
	Outcomes []int 
	ProbWeights []float64 
	CumProbs []float64
}

func NewOffspringDistProb(dist map[int]float64) *OffspringProb {
	
}
// simulation of a branching process
func BranchingMain() {
	// define the offspring distribution
	dist := map[int]float64 {
		0: 0.3,
		1: 0.4,
		2: 0.3,
	}

	offspringDist := NewOffspringDistProb(dist)

	// simulate the process
	result := simulateGW(offsprintDist, 100)

	// print the results:
	fmt.Println("Galton-Watson Branching process")
	for gen, size := range result {
		fmt.Printf("Generation %d: %d individuals\n", gen, size)
	}

}

func NewOffspringDistProb(dist map[int]float64) {

}