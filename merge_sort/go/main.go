package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 15

// Merge two sorted subarrays `arr[low … mid]` and `arr[mid+1 … high]`
func Merge(arr []int, aux []int, low int, mid int, high int) {
	k, i, j := low, low, mid+1

	// while there are elements in the left and right runs
	for i <= mid && j <= high {
		if arr[i] <= arr[j] {
			aux[k] = arr[i]
			k++
			i++
		} else {
			aux[k] = arr[j]
			k++
			j++
		}
	}
	// copy remaining elements
	for i <= mid {
		aux[k] = arr[i]
		k++
		i++
	}
	// No need to copy the second half (since the remaining items
	// are already in their correct position in the auxiliary array)

	// copy back to the original array to reflect sorted order
	for i := low; i <= high; i++ {
		arr[i] = aux[i]
	}
}

// Sort array `arr[low…high]` using auxiliary array `aux`
func mergesort(arr []int, aux []int, low int, high int) {

	// base case
	if high <= low {
		return
	}

	// find midpoint
	mid := low + ((high - low) >> 1)

	// recursively split runs into two halves until run size <= 1,
	// then merge them and return up the call chain

	mergesort(arr, aux, low, mid)    // split/merge left half
	mergesort(arr, aux, mid+1, high) // split/merge right half

	Merge(arr, aux, low, mid, high) // merge the two half runs.
}

// Function to check if arr is sorted in ascending order or not
func isSorted(arr []int) bool {
	for i := 1; i < N; i++ {
		if arr[i-1] > arr[i] {
			fmt.Printf("MergeSort Fails")
			return false
		}
	}
	return true
}

// Implementation of merge sort algorithm in go
func main() {
	var arr [N]int
	var aux [N]int
	rand.NewSource(time.Now().UnixNano())
	//rand.Seed(time.Now().UnixNano())

	// generate random input of integers
	for i := 0; i < N; i++ {
		ramdomValue := rand.Intn(100) - 50
		arr[i] = ramdomValue
		aux[i] = arr[i]
	}

	// sort array `arr` using auxiliary array `aux`
	mergesort(arr[:], aux[:], 0, N-1)

	if isSorted(arr[:]) {
		for i := 0; i < N; i++ {
			fmt.Printf("%d", arr[i])
		}
		fmt.Println()
	}

}
