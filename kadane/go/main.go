package main

import (
	"fmt"
)

// Function to find the maximum sum of a contiguous subarray
// in a given integer array
func kadane(arr []int) int {

	// find the maximum element present in a given array
	maxNum := arr[0]
	for _, num := range arr {
		if num > maxNum {
			maxNum = num
		}
	}

	// stores the maximum sum subarray found so far
	maxSoFar := 0

	// stores the maximum sum of subarray ending at the current position
	maxEndingHere := 0

	// traverse the given array
	for _, num := range arr {
		// update the maximum sum of subarray "ending" at index 'i' (by adding the
		// current element to maximum sum ending at previous index 'i-1')
		maxEndingHere = maxEndingHere + num

		//if the maximum sum is negative, set it to 0 (which represents
		// an empty subarray)
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}

		// update the result if the current subarray sum is found to be greater
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}

	}
	if maxSoFar == 0 {
		return maxNum
	}
	return maxSoFar
}

func main() {
	arr := []int{-8, -3, -6, -2, -5, -4}

	fmt.Printf("The maximum sum of a contiguous subarray is %d", kadane(arr))
}
