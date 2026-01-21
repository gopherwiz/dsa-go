package main

import "fmt"

func insertIterative(root *Node, val int) *Node {
	// If the tree is empty, the new node becomes the root
	if root == nil {
		return &Node{Val: val}
	}

	curr := root
	for {
		if val < curr.Val {
			// If target is smaller, we want to go Left
			if curr.Left == nil {
				curr.Left = &Node{Val: val}
				return root // Return from this branch
			}
			curr = curr.Left
		} else if val > curr.Val {
			// If target is larger, we want to go Right
			if curr.Right == nil {
				curr.Right = &Node{Val: val}
				return root // Return from this branch
			}
			curr = curr.Right
		} else {
			// Value already exists in the tree
			return root
		}
	}
}

func main() {
	var root *Node
	values := []int{50, 30, 20, 40, 70, 60, 80}

	for _, v := range values {
		root = insertIterative(root, v)
	}

	fmt.Print("Inorder Traversal: ")
	inorder(root)
}
