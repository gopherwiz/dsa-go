package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// findMin finds the smallest node (leftmost) in a subtree
func findMin(root *Node) *Node {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// deleteNode removes the first occurrence of target and returns the new root
func deleteNode(root *Node, target int) *Node {
	if root == nil {
		return nil
	}

	// 1. Locate the node
	if target < root.Val {
		root.Left = deleteNode(root.Left, target)
	} else if target > root.Val {
		root.Right = deleteNode(root.Right, target)
	} else {
		// Found the node to delete!

		// Case 1 & 2: No child or only one child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Case 3: Two children
		// Find the smallest in the right subtree (successor)
		successor := findMin(root.Right)
		// Replace current node's value with successor's value
		root.Val = successor.Val
		// Delete the successor node from the right subtree
		root.Right = deleteNode(root.Right, successor.Val)
	}

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

func main() {
	/*
	   Constructing BST:
	         50
	       /    \
	      30     70
	     /  \   /  \
	    20  40 60  80
	*/
	root := &Node{Val: 50}
	root.Left = &Node{Val: 30, Left: &Node{Val: 20}, Right: &Node{Val: 40}}
	root.Right = &Node{Val: 70, Left: &Node{Val: 60}, Right: &Node{Val: 80}}

	fmt.Print("Before deletion: ")
	inorder(root)
	fmt.Println()

	target := 30
	root = deleteNode(root, target)

	fmt.Printf("After deleting %d: ", target)
	inorder(root)
	fmt.Println()
}
