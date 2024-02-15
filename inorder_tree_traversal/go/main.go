package main

import "fmt"

// Node struct to store a binary tree node
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Constructor function to create a new Node
func newNode(data int) *Node {
	return &Node{data, nil, nil}
}

// Recursive function to perform inorder traversal on the tree
func inorder(root *Node) {
	// return if the current node is empty
	if root == nil {
		return
	}

	// Traverse the left subtree
	inorder(root.left)

	// Display the data part of the root (or current node)
	fmt.Printf("%d ", root.data)

	// Traverse the right subtree
	inorder(root.right)
}

func main() {
	/* Construct the following tree
	          1
	        /   \
	       /     \
	      2       3
	     /      /   \
	    /      /     \
	   4      5       6
	         / \
	        /   \
	       7     8
	*/

	root := newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.left.left = newNode(4)
	root.right.left = newNode(5)
	root.right.right = newNode(6)
	root.right.left.left = newNode(7)
	root.right.left.right = newNode(8)

	inorder(root)
}
