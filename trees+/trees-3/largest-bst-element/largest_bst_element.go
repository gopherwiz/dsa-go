package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func findMax(root *Node) *Node {
	if root == nil {
		return nil
	}

	curr := root
	// Keep moving to the right child as long as it exists
	for curr.Right != nil {
		curr = curr.Right
	}

	return curr
}

func main() {
	// Constructing BST:
	//      10
	//     /  \
	//    5    20
	//          \
	//           30  <-- Largest
	root := &Node{Val: 10}
	root.Left = &Node{Val: 5}
	root.Right = &Node{Val: 20}
	root.Right.Right = &Node{Val: 30}

	maxNode := findMax(root)

	if maxNode != nil {
		fmt.Printf("The largest element is: %d\n", maxNode.Val)
	} else {
		fmt.Println("The tree is empty.")
	}
}
