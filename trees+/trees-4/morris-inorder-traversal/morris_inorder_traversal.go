package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func morrisInorder(root *TreeNode) []int {
	var result []int
	curr := root

	for curr != nil {
		if curr.Left == nil {
			// No left child: visit and move right
			result = append(result, curr.Val)
			curr = curr.Right
		} else {
			// Find the inorder predecessor
			pre := curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}

			if pre.Right == nil {
				// Create a temporary thread to the successor (curr)
				pre.Right = curr
				curr = curr.Left
			} else {
				// Thread already exists: we are done with left subtree
				pre.Right = nil // Restore the tree
				result = append(result, curr.Val)
				curr = curr.Right
			}
		}
	}

	return result
}

func main() {
	/*
	      4
	     / \
	    2   6
	   / \
	  1   3
	*/
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	root.Right = &TreeNode{Val: 6}

	fmt.Printf("Morris Inorder: %v\n", morrisInorder(root))
	// Output: [1 2 3 4 6]
}

func morrisPreorder(root *TreeNode) []int {
	var result []int
	curr := root

	for curr != nil {
		if curr.Left == nil {
			result = append(result, curr.Val) // Visit root
			curr = curr.Right
		} else {
			pre := curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}

			if pre.Right == nil {
				// FIRST ENCOUNTER: Visit before going left
				result = append(result, curr.Val)
				pre.Right = curr
				curr = curr.Left
			} else {
				// SECOND ENCOUNTER: Just cleaning up
				pre.Right = nil
				curr = curr.Right
			}
		}
	}
	return result
}
