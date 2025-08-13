package main

import "fmt"

// main generates and prints Pascal's triangle
func main() {
	// Number of rows to generate in Pascal's triangle
	n := 10
	
	// Initialize triangle with the first row containing just 1
	triangle := [][]int{{1}}

	// Generate each row of Pascal's triangle
	for i := 1; i < n; i++ {
		// Create a new row with length i+1 (row number + 1)
		row := make([]int, i+1)
		
		// Set the first and last elements of each row to 1
		row[0], row[i] = 1, 1
		
		// Calculate the middle elements by summing the two elements above
		// Each element is the sum of the two elements from the previous row
		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		
		// Add the completed row to the triangle
		triangle = append(triangle, row)
	}

	// Print each row of the triangle
	for _, row := range triangle {
		fmt.Println(row)
	}
}