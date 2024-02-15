package main

import "fmt"

// Data structure to store a binary tree node
type Node struct {
	data  int
	left  *Node
	right *Node
}

//function to perform preorder traversal on the tree
func preorder(root *Node) {

	// if the current node is empty
	if root == nil {
		return
	}

	// Display the data part of the root (or current node)
	fmt.Printf("%d ", root.data)

	// Traverse the left subtree
	preorder(root.left)

	// Traverse the right subtree
	preorder(root.right)
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

	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.right.left = &Node{data: 5}
	root.right.right = &Node{data: 6}
	root.right.left.left = &Node{data: 7}
	root.right.left.right = &Node{data: 8}

	preorder(root)

}
