package main

import (
	"fmt"
)

// Data structure to store a graph edge
type Edge struct {
	src, dest int
}

// A structure to represent a graph object
type Graph struct {
	// a slice of slices to represent an adjacency list
	adjList [][]int
}

// Graph Constructor
func NewGraph(edges []Edge, n int) *Graph {
	// create a new graph object
	graph := &Graph{}

	// resize the adjacency list to hold `n` elements of type `[]int`
	graph.adjList = make([][]int, n)

	// add edges to the undirected graph
	for _, edge := range edges {
		graph.adjList[edge.src] = append(graph.adjList[edge.src], edge.dest)
		graph.adjList[edge.dest] = append(graph.adjList[edge.dest], edge.src)
	}

	return graph
}

// Perform BFS on the graph starting from vertex `v`
func BFS(graph *Graph, v int, discovered []bool) {
	// create a queue for doing BFS
	q := make([]int, 0)

	// mark the source vertex as discovered
	discovered[v] = true

	// enqueue source vertex
	q = append(q, v)

	// loop till queue is empty
	for len(q) > 0 {
		// dequeue front node and print it
		v := q[0]
		q = q[1:]
		fmt.Printf("%d ", v)

		// do for every edge (v, u)
		for _, u := range graph.adjList[v] {
			if !discovered[u] {
				// mark it as discovered and enqueue it
				discovered[u] = true
				q = append(q, u)
			}
		}
	}
}

func main() {
	// vector of graph edges as per the above diagram
	edges := []Edge{
		{1, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}, {5, 9},
		{5, 10}, {4, 7}, {4, 8}, {7, 11}, {7, 12},
		// vertex 0, 13, and 14 are single nodes
	}

	// total number of nodes in the graph (labelled from 0 to 14)
	n := 15

	// build a graph from the given edges
	graph := NewGraph(edges, n)

	// to keep track of whether a vertex is discovered or not
	discovered := make([]bool, n)

	// Perform BFS traversal from all undiscovered nodes to
	// cover all connected components of a graph
	for i := 0; i < n; i++ {
		if !discovered[i] {
			// start BFS traversal from vertex `i`
			BFS(graph, i, discovered)
		}
	}
}
