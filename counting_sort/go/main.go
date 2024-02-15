package main

import "fmt"

// Function to perform count sort
func countsort(arr []int, n, k int) {

	// Create a vector to store the sorted array
	output := make([]int, n)

	// Create a vector to store frequency of each element, initialized with zeros
	freq := make([]int, k+1)

	// 1. Count the frequency of each element in the input array
	for i := 0; i < n; i++ {
		freq[arr[i]]++
	}

	// 2. Calculate cumulative frequency
	total := 0
	for i := 0; i < k+1; i++ {
		oldCount := freq[i]
		freq[i] = total
		total += oldCount
	}

	// 3. Place elements in correct position in the output array
	for i := 0; i < n; i++ {
		output[freq[arr[i]]] = arr[i]
		freq[arr[i]]++
	}

	// Copy the sorted elements back to the original array
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func main() {
	// Input array
	arr := []int{4, 2, 10, 10, 1, 4, 2, 1, 10}

	// Calculate the size of the array
	n := len(arr)

	// Range of array elements
	k := 10

	// Perform count sort
	countsort(arr, n, k)

	// Print the sorted array
	for i := 0; i < n; i++ {
		fmt.Printf("%d", arr[i])
	}
	fmt.Println()
}
