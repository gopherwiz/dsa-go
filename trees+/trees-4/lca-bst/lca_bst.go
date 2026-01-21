package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// FindLCABST takes advantage of BST properties.
// Time: O(h), Space: O(1)
func FindLCABST(root *Node, p int, q int) int {
	curr := root

	for curr != nil {
		// If both target values are smaller than current, move left
		if p < curr.Val && q < curr.Val {
			curr = curr.Left
		} else if p > curr.Val && q > curr.Val {
			// If both target values are larger than current, move right
			curr = curr.Right
		} else {
			// We found the "split" point.
			// This happens when p <= curr.Val <= q (or vice versa).
			return curr.Val
		}
	}

	return -1
}

func main() {
	/*
	   Constructing the BST:
	           20
	          /  \
	         10   30
	        /  \
	       5    15
	      /
	     2
	*/
	root := &Node{Val: 20}
	root.Left = &Node{Val: 10}
	root.Right = &Node{Val: 30}
	root.Left.Left = &Node{Val: 5}
	root.Left.Right = &Node{Val: 15}
	root.Left.Left.Left = &Node{Val: 2}

	// Example 1: LCA of 2 and 15
	// 10 is the split point (2 < 10 and 15 > 10)
	fmt.Printf("LCA of 2 and 15: %d\n", FindLCABST(root, 2, 15))

	// Example 2: LCA of 5 and 15
	// 10 is the split point
	fmt.Printf("LCA of 5 and 15: %d\n", FindLCABST(root, 5, 15))

	// Example 3: LCA of 2 and 5
	// 5 is the LCA because 5 is the ancestor of 2 and 5 is one of the targets
	fmt.Printf("LCA of 2 and 5: %d\n", FindLCABST(root, 2, 5))

	// Example 4: LCA of 15 and 30
	// 20 is the root/split point
	fmt.Printf("LCA of 15 and 30: %d\n", FindLCABST(root, 15, 30))
}
