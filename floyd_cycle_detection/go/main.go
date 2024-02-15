package main

import (
	"fmt"
)

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

func detectCycle(head *Node) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}
func main() {
	// input keys
	keys := []int{1, 2, 3, 4, 5}

	var head *Node

	for i := len(keys) - 1; i >= 0; i-- {
		push(&head, keys[i])
	}

	// insert cycle
	head.next.next.next.next.next = head.next.next

	if detectCycle(head) {
		fmt.Println("Cycle Found")
	} else {
		fmt.Println("No Cycle Found")
	}
}
