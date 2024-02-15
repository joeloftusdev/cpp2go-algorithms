package main

import (
	"fmt"
	"math"
)

// Recursive function to print path of given vertex `u` from source vertex `v`
func printPath(path [][]int, v, u int) {
	if path[v][u] == v {
		return
	}
	printPath(path, v, path[v][u])
	fmt.Printf("%d, ", path[v][u])
}

// Function to print the shortest cost with path information between
// all pairs of vertices

func printSolution(cost, path [][]int) {
	n := len(cost)
	for v := 0; v < n; v++ {
		for u := 0; u < n; u++ {
			if u != v && path[v][u] != -1 {
				fmt.Printf("The shortest path from %d —> %d is [%d, ", v, u, v)
				printPath(path, v, u)
				fmt.Printf("%d]\n", u)
			}
		}
	}
}

// Function to run the Floyd–Warshall algorithm
func floydWarshall(adjMatrix [][]int) {

	// total number of vertices in the `adjMatrix`
	n := len(adjMatrix)

	// base case
	if n == 0 {
		return
	}

	// cost[] and path[] stores shortest path
	// (shortest cost/shortest route) information
	cost := make([][]int, n)
	path := make([][]int, n)

	// initialize cost[] and path[]
	for v := 0; v < n; v++ {
		cost[v] = make([]int, n)
		path[v] = make([]int, n)
		for u := 0; u < n; u++ {
			// initially, cost would be the same as the weight of the edge
			cost[v][u] = adjMatrix[v][u]

			if v == u {
				path[v][u] = 0
			} else if cost[v][u] != math.MaxInt64 {
				path[v][u] = v
			} else {
				path[v][u] = -1
			}
		}
	}
	// run Floyd–Warshall
	for k := 0; k < n; k++ {
		for v := 0; v < n; v++ {
			for u := 0; u < n; u++ {
				// If vertex `k` is on the shortest path from `v` to `u`,
				// then update the value of cost[v][u] and path[v][u]

				if cost[v][k] != math.MaxInt64 && cost[k][u] != math.MaxInt64 &&
					cost[v][k]+cost[k][u] < cost[v][u] {
					cost[v][u] = cost[v][k] + cost[k][u]
					path[v][u] = path[k][u]
				}
			}
			// if diagonal elements become negative, the
			// graph contains a negative-weight cycle
			if cost[v][v] < 0 {
				fmt.Printf("Negative-weight cycle found!!")
				return
			}
		}
	}
	// Print the shortest path between all pairs of vertices
	printSolution(cost, path)
}

func main() {
	// define infinity
	I := math.MaxInt64

	// given adjacency representation of the matrix
	adjMatrix := [][]int{
		{0, I, -2, I},
		{4, 0, 3, I},
		{I, I, 0, 2},
		{I, -1, I, 0},
	}

	// Run Floyd–Warshall algorithm
	floydWarshall(adjMatrix)

}
