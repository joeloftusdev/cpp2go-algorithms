package main

import (
	"fmt"
	"math"
)

// Data structure to store a graph edge
type Edge struct {
	source, dest, weight int
}

// Recursive function to print the path of a given vertex from source vertex
func printPath(parent []int, vertex, source int) {

	if vertex == source || vertex < 0 {
		fmt.Printf("%d", vertex)
		return
	}
	printPath(parent, parent[vertex], source)

	fmt.Printf(", %d", vertex)
}

// Function to run the Bellman–Ford algorithm from a given source
func bellmanFord(edges []Edge, source, n int) {
	// distance[] and parent[] stores the shortest path (least cost/path)
	// information. Initially, all vertices except the source vertex
	// weight INFINITY and no parent

	distance := make([]int, n)
	parent := make([]int, n)
	for i := range distance {
		distance[i] = math.MaxInt64
	}
	distance[source] = 0

	for k := 0; k < n-1; k++ {
		for _, edge := range edges {
			u, v, w := edge.source, edge.dest, edge.weight
			// if the distance to destination `v` can be
			// shortened by taking edge (u, v)
			if distance[u] != math.MaxInt64 && distance[u]+w < distance[v] {
				// update distance to the new lower value
				distance[v] = distance[u] + w
				// set v's parent as `u`
				parent[v] = u
			}
		}
	}

	// check for negative-weight cycles
	for _, edge := range edges {
		// edge from `u` to `v` having weight `w`
		u, v, w := edge.source, edge.dest, edge.weight

		// if the distance to destination `u` can be shortened by taking edge (u, v)
		if distance[u] != math.MaxInt64 && distance[u]+w < distance[v] {
			fmt.Printf("Negative-weight cycle is found!!")
			return
		}
	}

	for i := 0; i < n; i++ {
		if i != source && distance[i] < math.MaxInt64 {
			fmt.Printf("The distance of vertex %d from the source is %2d. Its path is [", i, distance[i])
			printPath(parent, i, source)
			fmt.Println("]")
		}
	}

}

func main() {
	edges := []Edge{
		// (x, y, w) —> edge from `x` to `y` having weight `w`
		{0, 1, -1}, {0, 2, 4}, {1, 2, 3}, {1, 3, 2},
		{1, 4, 2}, {3, 2, 5}, {3, 1, 1}, {4, 3, -3},
	}

	// set the maximum number of nodes in the graph
	n := 5

	// run the Bellman–Ford algorithm from every node
	for source := 0; source < n; source++ {
		bellmanFord(edges, source, n)
	}
}
