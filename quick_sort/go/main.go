package main

import "fmt"

// Partition using the Lomuto partition scheme
func partition(a []int, start int, end int) int {

	// Pick the rightmost element as a pivot from the array
	pivot := a[end]

	// elements less than the pivot will be pushed to the left of `pIndex`
	// elements more than the pivot will be pushed to the right of `pIndex`
	// equal elements can go either way

	pIndex := start

	// each time we find an element less than or equal to the pivot, `pIndex`
	// is incremented, and that element would be placed before the pivot.

	for i := start; i < end; i++ {
		if a[i] <= pivot {
			a[i], a[pIndex] = a[pIndex], a[i]
			pIndex++
		}
	}

	// swap `pIndex` with pivot
	a[pIndex], a[end] = a[end], a[pIndex]

	// return `pIndex` (index of the pivot element)
	return pIndex

}

func quicksort(a []int, start int, end int) {
	// base condition
	if start >= end {
		return
	}

	// rearrange elements across pivot
	pivot := partition(a, start, end)

	// recur on subarray containing elements that are less than the pivot
	quicksort(a, start, pivot-1)

	// recur on subarray containing elements that are more than the pivot
	quicksort(a, pivot+1, end)

}

// Golang implementation of the Quicksort algorithm

func main() {
	a := []int{9, -3, 5, 2, 6, 8, -6, 1, 3}
	n := len(a)

	quicksort(a, 0, n-1)

	// print the sorted array
	for i := 0; i < n; i++ {
		fmt.Printf("%d", a[i])
	}

}
