package main

import "fmt"

// A Linked List Node
type Node struct {
	data int
	next *Node
}

// Helper function to create a new node with the given data and
// pushes it onto the list's front

func push(headRef **Node, data int) {
	newNode := &Node{
		data: data,
		next: *headRef,
	}
	*headRef = newNode
}

// Function to detect a cycle in a linked list using hashing
func detectCycle(head *Node) bool {
	curr := head
	set := make(map[*Node]struct{})

	// traverse the list

	for curr != nil {
		// If the current node is already in the set the cycle is detected
		if _, found := set[curr]; found {
			return true
		}
		set[curr] = struct{}{}

		curr = curr.next

	}

	return false
}

func main() {
	keys := []int{1, 2, 3, 4, 5}

	var head *Node
	for i := len(keys) - 1; i >= 0; i-- {
		push(&head, keys[i])
	}

	// Insert cycle
	head.next.next.next.next.next = head.next.next

	// Detect cycle and print result
	if detectCycle(head) {
		fmt.Println("Cycle Found")
	} else {
		fmt.Println("No Cycle Found")
	}
}
