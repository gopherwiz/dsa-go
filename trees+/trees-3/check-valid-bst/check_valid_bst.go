package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) int {
	// Start with the widest possible range
	if checkValid(root, math.MinInt64, math.MaxInt64) {
		return 1
	}
	return 0
}

func checkValid(root *TreeNode, min int64, max int64) bool {
	if root == nil {
		return true
	}

	val := int64(root.Val)

	// 1. Check if current node violates the range constraints
	if val <= min || val >= max {
		return false
	}

	// 2. Recursively check subtrees with updated constraints
	// Left subtree nodes must be < current root.Val
	// Right subtree nodes must be > current root.Val
	return checkValid(root.Left, min, val) &&
		checkValid(root.Right, val, max)
}

func main() {
	// Valid BST
	//      20
	//     /  \
	//    10   30
	root := &TreeNode{Val: 20}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 30}

	fmt.Printf("Is Valid BST: %d\n", isValidBST(root)) // Output: 1
}
