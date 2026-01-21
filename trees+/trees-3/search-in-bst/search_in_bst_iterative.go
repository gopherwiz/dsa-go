package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// searchIterative uses a loop and O(1) extra space.
func searchIterative(root *Node, target int) *Node {
	curr := root
	for curr != nil {
		if curr.Val == target {
			return curr
		}
		if target < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}

func main() {
	/*
	   Constructing this BST:
	         10
	       /    \
	      5      15
	     / \    /  \
	    2   7  12   20
	*/
	root := &Node{Val: 10}
	root.Left = &Node{Val: 5, Left: &Node{Val: 2}, Right: &Node{Val: 7}}
	root.Right = &Node{Val: 15, Left: &Node{Val: 12}, Right: &Node{Val: 20}}

	targets := []int{7, 12, 100}

	fmt.Println("Searching BST:")
	for _, t := range targets {
		// Using the iterative search
		result := searchIterative(root, t)

		if result != nil {
			fmt.Printf("Target %d: FOUND at node address %p\n", t, result)
		} else {
			fmt.Printf("Target %d: NOT FOUND\n", t)
		}
	}
}
