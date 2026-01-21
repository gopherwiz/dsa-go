package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func search(root *Node, target int) *Node {
	// Base Case: Not found or empty tree
	if root == nil {
		return nil
	}

	// Found it!
	if root.Val == target {
		return root
	}

	// Search Left
	leftRes := search(root.Left, target)
	if leftRes != nil {
		return leftRes
	}

	// If not in left, search Right
	return search(root.Right, target)
}

func main() {
	/* Tree:
	       1
	      / \
	     2   3
	    / \
	   4   5
	*/
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}
	root.Right = &Node{Val: 3}

	target := 5
	result := search(root, target)

	if result != nil {
		fmt.Printf("Found node %d at memory address %p\n", target, result)
	} else {
		fmt.Println("Value not found")
	}
}

func searchBFS(root *Node, target int) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Val == target {
			return curr
		}

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return nil
}
