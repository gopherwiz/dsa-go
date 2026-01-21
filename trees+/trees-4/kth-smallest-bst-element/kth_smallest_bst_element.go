package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	result := -1

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		// Optimization: if result is found or node is nil, stop recursing
		if node == nil || result != -1 {
			return
		}

		// 1. Traverse Left
		traverse(node.Left)

		// 2. Visit (Check Counter)
		count++
		if count == k {
			result = node.Val
			return
		}

		// 3. Traverse Right
		traverse(node.Right)
	}

	traverse(root)
	return result
}

func main() {
	/*
		      5
		     / \
		    3   6
		   / \
		  2   4
		 /
		1
	*/
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}
	root.Right = &TreeNode{Val: 6}

	k := 3
	fmt.Printf("The %d-rd smallest element is: %d\n", k, kthSmallest(root, k))
	// Output: 3 (Nodes visited: 1, 2, 3...)
}
