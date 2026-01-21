package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// insert returns the root of the tree after insertion
func insert(root *Node, val int) *Node {
	// Base case: if we found the empty spot, create the node
	if root == nil {
		return &Node{Val: val}
	}

	if val < root.Val {
		// If value is smaller, go left
		root.Left = insert(root.Left, val)
	} else if val > root.Val {
		// If value is larger, go right
		root.Right = insert(root.Right, val)
	}

	// Return the (unchanged) node pointer
	return root
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Printf("%d ", root.Val)
	inorder(root.Right)
}

//func main() {
//	var root *Node
//
//	// Creating the tree
//	values := []int{50, 30, 20, 40, 70, 60, 80}
//	for _, v := range values {
//		root = insert(root, v)
//	}
//
//	fmt.Print("Inorder Traversal (should be sorted): ")
//	inorder(root)
//	// Output: 20 30 40 50 60 70 80
//}
