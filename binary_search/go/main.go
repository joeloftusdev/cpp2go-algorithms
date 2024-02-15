package main

import (
	"fmt"
)

// Iterative implementation of the binary search algorithm to return
// the position of `target` in array `nums` of size `n`
func binarySearch(nums []int, target int) int {
	// search space is nums[low…high]
	low, high := 0, len(nums)-1

	// loop till the search space is exhausted
	for low <= high {
		// find the mid-value in the search space and
		// compares it with the target
		mid := low + (high-low)/2 // updated to avoid overflow

		// target value is found
		if target == nums[mid] {
			return mid
		}

		// if the target is less than the middle element, discard all elements
		// in the right search space, including the middle element
		if target < nums[mid] {
			high = mid - 1
		} else { // if the target is more than the middle element, discard all elements
			// in the left search space, including the middle element
			low = mid + 1
		}
	}

	// target doesn't exist in the array
	return -1
}

func main() {
	nums := []int{2, 5, 6, 8, 9, 10}
	target := 5

	index := binarySearch(nums, target)

	if index != -1 {
		fmt.Printf("Element found at index %d\n", index)
	} else {
		fmt.Println("Element not found in the array")
	}

}
