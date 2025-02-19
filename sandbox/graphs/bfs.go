package main

import (
	"container/list"
	"fmt"
)

// represent a graph as an adjacency list
// e.g.
// A -> {B, C}
// C -> {D}
// D -> {}
type Graph struct {
	vertices map[string][]string
}

// Create a new graph with an empty adjacency list
func NewGraph() *Graph {
	return &Graph{vertices: make(map[string][]string)}
}

// Add  an edge to the graph
func (g *Graph) AddEdge(v, w string) {
	g.vertices[v] = append(g.vertices[v], w)
	g.vertices[w] = append(g.vertices[w], v)
}

// Perform a breadth-first search on the graph starting from the given vertex
func (g *Graph) BFS(start string) {
	// make a map to keep track of visited vertices
	visited := make(map[string]bool)
	// use a queue to keep track of vertices to visit
	// use a doubly-linked list from  the package container/list
	queue := list.New()

	visited[start] = true
	queue.PushBack(start)

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(string)
		fmt.Printf("%s ", node)
		queue.Remove(element)

		for _, neighbour := range g.vertices[node] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue.PushBack(neighbour)
			}
		}
	}

}

func main() {
	g := NewGraph()
	g.AddEdge("start", "A")
	g.AddEdge("start", "B")
	g.AddEdge("A", "finish")
	g.AddEdge("B", "A")
	g.AddEdge("B", "finish")
	g.BFS("start")
	
}