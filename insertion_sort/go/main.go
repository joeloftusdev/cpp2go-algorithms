package main

import "fmt"

// Function to perform insertion sort on `arr[]`
func insertionSort(arr []int) {
	// start from the second element (the element at index 0
	// is already sorted)
	for i := 1; i < len(arr); i++ {

		value := arr[i]
		j := i

		// find index `j` within the sorted subset `arr[0…i-1]`
		// where element `arr[i]` belongs
		for j > 0 && arr[j-1] > value {
			arr[j] = arr[j-1]
			j--
		}

		// note that the subarray `arr[j…i-1]` is shifted to
		// the right by one position, i.e., `arr[j+1…i]`

		arr[j] = value
	}
}

// Function to print `n` elements of array `arr`
func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
	}
}

func main() {
	arr := []int{3, 8, 5, 4, 1, 9, -2}

	insertionSort(arr)

	// print the sorted array
	printArray(arr)
}
