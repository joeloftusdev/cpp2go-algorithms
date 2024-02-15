package main

import (
	"fmt"
)

// Below arrays detail all eight possible movements
var row = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var col = []int{-1, 0, 1, -1, 1, -1, 0, 1}

// check if it is possible to go to pixel (x, y) from the
// current pixel. The function returns false if the pixel
// has a different color, or it's not a valid pixel

func isSafe(mat [][]rune, x, y int, target rune) bool {
	return x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) && mat[x][y] == target
}

// Flood fill using DFS
func floodfill(mat [][]rune, x, y int, replacement rune) {
	// base case
	if len(mat) == 0 {
		return
	}
	// get the target color
	target := mat[x][y]

	// target color is same as replacement
	if target == replacement {
		return
	}

	// replace the current pixel color with that of replacement
	mat[x][y] = replacement

	// process all eight adjacent pixels of the current pixel and
	// recur for each valid pixel

	for k := 0; k < 8; k++ {
		// if the adjacent pixel at position (x + row[k], y + col[k]) is
		// a valid pixel and has the same color as that of the current pixel
		if isSafe(mat, x+row[k], y+col[k], target) {
			floodfill(mat, x+row[k], y+col[k], replacement)
		}
	}

}

// Utility function to print a matrix
func printMatrix(mat [][]rune) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Printf("%c ", mat[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// matrix showing portion of the screen having different colors
	mat := [][]rune{
		{'Y', 'Y', 'Y', 'G', 'G', 'G', 'G', 'G', 'G', 'G'},
		{'Y', 'Y', 'Y', 'Y', 'Y', 'Y', 'G', 'X', 'X', 'X'},
		{'G', 'G', 'G', 'G', 'G', 'G', 'G', 'X', 'X', 'X'},
		{'W', 'W', 'W', 'W', 'W', 'G', 'G', 'G', 'G', 'X'},
		{'W', 'R', 'R', 'R', 'R', 'R', 'G', 'X', 'X', 'X'},
		{'W', 'W', 'W', 'R', 'R', 'G', 'G', 'X', 'X', 'X'},
		{'W', 'B', 'W', 'R', 'R', 'R', 'R', 'R', 'R', 'X'},
		{'W', 'B', 'B', 'B', 'B', 'R', 'R', 'X', 'X', 'X'},
		{'W', 'B', 'B', 'X', 'B', 'B', 'B', 'B', 'X', 'X'},
		{'W', 'B', 'B', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
	}

	// start node
	x, y := 3, 9 // having a target color `X`

	// replacement color
	replacement := 'C'

	// replace the target color with a replacement color using DFS
	floodfill(mat, x, y, replacement)

	// print the colors after replacement
	printMatrix(mat)
}
