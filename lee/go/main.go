package main

import (
	"fmt"
	"math"
)

// A Queue Node
type Node struct {

	// (x, y) represents matrix cell coordinates, and
	// `dist` represents their minimum distance from the source
	x, y, dist int
}

var row = []int{-1, 0, 0, 1}
var col = []int{0, -1, 1, 0}

// Function to check if it is possible to go to position (row, col)
// from the current position. The function returns false if (row, col)
// is not a valid position or has a value 0 or already visited.

func isValid(mat [][]int, visited [][]bool, row, col int) bool {
	return row >= 0 && row < len(mat) && col >= 0 && col < len(mat[0]) && mat[row][col] == 1 && !visited[row][col]
}

// findShortestPathLength finds the shortest possible route in a matrix `mat` from source
// cell (i, j) to destination cell (x, y)
func findShortestPathLength(mat [][]int, src, dest []int) int {
	// Base case: invalid input
	if len(mat) == 0 || mat[src[0]][src[1]] == 0 || mat[dest[0]][dest[1]] == 0 {
		return -1
	}

	// Matrix dimensions
	M := len(mat)
	N := len(mat[0])

	// Construct a matrix to keep track of visited cells
	visited := make([][]bool, M)
	for i := range visited {
		visited[i] = make([]bool, N)
	}

	// Create an empty queue
	q := make([]Node, 0)

	// Get source cell coordinates
	i, j := src[0], src[1]

	// Mark the source cell as visited and enqueue the source node
	visited[i][j] = true
	q = append(q, Node{i, j, 0})

	// Store length of the shortest path from source to destination
	minDist := 1<<31 - 1 // Max int value in Go

	// Loop till queue is empty
	for len(q) > 0 {
		// Dequeue front node and process it
		node := q[0]
		q = q[1:]

		// Current cell coordinates and distance
		i, j, dist := node.x, node.y, node.dist

		// If destination is found, update `minDist` and stop
		if i == dest[0] && j == dest[1] {
			minDist = dist
			break
		}

		// Check for all possible movements from the current cell
		// and enqueue each valid movement
		for k := 0; k < 4; k++ {
			X := i + row[k]
			Y := j + col[k]
			// Check if it is possible to go to position
			// (i + row[k], j + col[k]) from current position
			if isValid(mat, visited, X, Y) {
				// Mark next cell as visited and enqueue it
				visited[X][Y] = true
				q = append(q, Node{X, Y, dist + 1})
			}
		}
	}

	if minDist != math.MaxInt64 {
		return minDist
	}

	return -1
}

func main() {
	mat := [][]int{
		{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 0, 1, 0, 1},
		{0, 0, 1, 0, 1, 1, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 0, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 1, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
		{0, 0, 1, 0, 0, 1, 1, 0, 0, 1},
	}

	src := []int{0, 0}
	dest := []int{7, 5}

	minDist := findShortestPathLength(mat, src, dest)

	if minDist != -1 {
		fmt.Println("The shortest path from source to destination has length:", minDist)
	} else {
		fmt.Println("Destination cannot be reached from the given source")
	}

}
