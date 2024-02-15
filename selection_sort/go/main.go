package main

import "fmt"

// Utility function to swap values at two indices in an array
func swap(arr []int, j int, i int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// Function to perform selection sort on `arr[]`
func selectionSort(arr []int) {
	// run `n-1` times
	for i := 0; i < len(arr)-1; i++ {

		// find the minimum element in the unsorted subarray `[i…n-1]`
		// and swap it with `arr[i]`
		min := i

		for j := i + 1; j < len(arr); j++ {
			// if `arr[j]` is less, then it is the new minimum
			if arr[j] < arr[min] {
				min = j // update the index of minimum element
			}
		}

		// // swap the minimum element in subarray `arr[i…n-1]` with `arr[i]`
		swap(arr, min, i)
	}
}

// Function to print `n` elements of array `arr`
func printArray(arr []int) {
	for _, i := range arr {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

func main() {
	arr := []int{3, 5, 8, 4, 1, 9, -2}

	selectionSort(arr)
	printArray(arr)
}
