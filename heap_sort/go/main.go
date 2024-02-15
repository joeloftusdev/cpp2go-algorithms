package main

import "fmt"

type PriorityQueue struct {
}

// return left child of `A[i]`
func (pq *PriorityQueue) left(i int) int {
	return 2*i + 1
}

// return right child of `A[i]`
func (pq *PriorityQueue) right(i int) int {
	return 2*i + 2
}

// Recursive heapify-down algorithm
// the node at index `i` and its two direct children
// violates the heap property

func (pq *PriorityQueue) heapify(A []int, i, size int) {
	// get left and right child of node at index `i`
	left := pq.left(i)
	right := pq.right(i)

	largest := i

	// compare `A[i]` with its left and right child
	// and find the largest value

	if left < size && A[left] > A[largest] {
		largest = left
	}

	if right < size && A[right] > A[largest] {
		largest = right
	}

	// swap with a child having greater value and
	// call heapify-down on the child
	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		pq.heapify(A, largest, size)
	}
}

// Constructor (Build-Heap)
func NewPriorityQueue(A []int, n int) *PriorityQueue {
	pq := &PriorityQueue{}

	// call heapify starting from the last internal node all the
	// way up to the root node

	i := (n - 2) / 2
	for i >= 0 {
		pq.heapify(A, i, n)
		i--
	}
	return pq
}

// Function to remove an element with the highest priority (present at the root)
func (pq *PriorityQueue) pop(A []int, size int) int {
	// if the heap has no elements
	if size <= 0 {
		return -1
	}

	top := A[0]

	// replace the root of the heap with the last element
	// of the array
	A[0] = A[size-1]

	// call heapify-down on the root node
	pq.heapify(A, 0, size-1)

	return top
}

// Function to perform heapsort on array `A` of size `n`
func heapsort(A []int, n int) {
	// build a priority queue and initialize it by the given array
	pq := NewPriorityQueue(A, n)

	// repeatedly pop from the heap till it becomes empty
	for n > 0 {
		A[n-1] = pq.pop(A, n)
		n--
	}

}

// Heapsort algorithm implementation in Go
func main() {
	A := []int{6, 4, 7, 1, 9, -2}
	n := len(A)

	// perform heapsort on the array
	heapsort(A, n)

	// print the sorted array
	for i := 0; i < n; i++ {
		fmt.Printf("%d", A[i])
	}
}
