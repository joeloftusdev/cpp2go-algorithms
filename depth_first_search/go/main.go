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

// Function to perform DFS traversal on the graph on a graph
func DFS(graph *Graph, v int, discovered []bool) {
	// mark the current node as discovered
	discovered[v] = true

	// print the current node
	fmt.Printf("%d ", v)

	// do for every edge (v, u)
	for _, u := range graph.adjList[v] {
		// if `u` is not yet discovered
		if !discovered[u] {
			DFS(graph, u, discovered)
		}
	}
}

func main() {
	// vector of graph edges as per the above diagram
	edges := []Edge{
		// Notice that node 0 is unconnected
		{1, 2}, {1, 7}, {1, 8}, {2, 3}, {2, 6}, {3, 4},
		{3, 5}, {8, 9}, {8, 12}, {9, 10}, {9, 11},
	}

	// total number of nodes in the graph (labelled from 0 to 12)
	n := 13

	// build a graph from the given edges
	graph := NewGraph(edges, n)

	// to keep track of whether a vertex is discovered or not
	discovered := make([]bool, n)

	// Perform DFS traversal from all undiscovered nodes to
	// cover all connected components of a graph
	for i := 0; i < n; i++ {
		if !discovered[i] {
			DFS(graph, i, discovered)
		}
	}
}
