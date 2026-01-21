package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Returns the height if balanced, otherwise returns -1
func getBalancedHeight(root *TreeNode) int {
	if root == nil {
		return 0 // Standard height for empty node
	}

	// 1. Get height of left subtree
	leftH := getBalancedHeight(root.Left)
	if leftH == -1 {
		return -1
	} // Already unbalanced

	// 2. Get height of right subtree
	rightH := getBalancedHeight(root.Right)
	if rightH == -1 {
		return -1
	} // Already unbalanced

	// 3. Check current node balance
	diff := leftH - rightH
	if diff < -1 || diff > 1 {
		return -1 // Unbalanced
	}

	// 4. Return actual height
	if leftH > rightH {
		return leftH + 1
	}
	return rightH + 1
}

func main() {
	// Example: Balanced Tree (Height 2)
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	height := getBalancedHeight(root)
	if height == -1 {
		fmt.Println("Tree is not balanced")
	} else {
		fmt.Printf("Tree is balanced with height: %d\n", height)
	}
}
