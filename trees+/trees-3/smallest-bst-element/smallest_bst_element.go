package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// FindSmallest recursively traverses to the leftmost node
func FindSmallest(root *Node) *Node {
	// 1. Handle empty tree
	if root == nil {
		return nil
	}

	// 2. Base case: No left child means this is the smallest value
	if root.Left == nil {
		return root
	}

	// 3. Recursive step: Move to the left child
	return FindSmallest(root.Left)
}

func main() {
	/*
	   Manually building (Hardcoding) this BST:
	             100
	            /   \
	           50    150
	          /  \
	         25   75
	        /
	       12  <-- Smallest
	*/

	root := &Node{Val: 100}
	root.Left = &Node{Val: 50}
	root.Right = &Node{Val: 150}

	root.Left.Left = &Node{Val: 25}
	root.Left.Right = &Node{Val: 75}

	root.Left.Left.Left = &Node{Val: 12}

	// Execute search
	smallest := FindSmallest(root)

	if smallest != nil {
		fmt.Printf("Smallest value found: %d\n", smallest.Val)
	} else {
		fmt.Println("The tree is empty.")
	}
}
