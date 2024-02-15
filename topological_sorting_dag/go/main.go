package main

import "fmt"

// Data structure to store a graph edge
type Edge struct {
	src, dest int
}

// A struct to represent a graph object

type Graph struct {
	// a slice of slices to represent adjList
	adjList [][]int
}

// Graph Constructor
func NewGraph(edges []Edge, n int) *Graph {
	graph := &Graph{}
	// resize the vector to hold `n` elements of type `<int>`
	graph.adjList = make([][]int, n)

	// add edges to the directed graph
	for _, edge := range edges {
		graph.adjList[edge.src] = append(graph.adjList[edge.src], edge.dest)
	}
	return graph
}

// Perform DFS on the graph and set the departure time of all
// vertices of the graph
func DFS(graph *Graph, v int, discovered []bool, departure []int, time *int) {
	// mark the current node as discovered
	discovered[v] = true
	// set the arrival time of `v`
	//*time++

	// do for every edge (v, u)
	for _, u := range graph.adjList[v] {
		// if `u` is not yet discovered
		if !discovered[u] {
			DFS(graph, u, discovered, departure, time)
		}
	}
	// ready to backtrack
	// set departure time of vertex `v`
	departure[*time] = v
	*time++

}

// Function to perform a topological sort on a given DAG
func doTopologicalSort(graph *Graph, n int) {
	// departure[] stores the vertex number using departure time as an index
	departure := make([]int, n)

	/* If we had done it the other way around, i.e., fill the array
	   with departure time using vertex number as an index, we would
	   need to sort it later */

	// to keep track of whether a vertex is discovered or not
	discovered := make([]bool, n)
	time := 0

	// perform DFS on all undiscovered vertices
	for i := 0; i < n; i++ {
		if !discovered[i] {
			DFS(graph, i, discovered, departure, &time)
		}
	}

	// Print the vertices in order of their decreasing
	// departure time in DFS, i.e., in topological order
	for i := n - 1; i >= 0; i-- {
		if departure[i] != -1 {
			fmt.Printf("%d ", departure[i])
		}
	}
}

func main() {
	// slice of graph edges as per the above diagram

	edges := []Edge{
		{0, 6}, {1, 2}, {1, 4}, {1, 6}, {3, 0}, {3, 4}, {5, 1}, {7, 0}, {7, 1},
	}

	// total number of nodes in the graph (labelled from 0 to 7)
	n := 8

	// build a graph from the given edges
	graph := NewGraph(edges, n)

	// perform topological sort
	doTopologicalSort(graph, n)
}
