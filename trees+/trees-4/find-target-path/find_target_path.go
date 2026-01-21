package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// GetPath returns the path from root to target as a slice.
// If target is not found, it returns an empty slice.
func GetPath(root *Node, target int) []int {
	var path []int
	if isPath := solve(root, target, &path); !isPath {
		return nil
	}

	reverse(path)

	return path
}

func solve(curr *Node, target int, path *[]int) bool {
	if curr == nil {
		return false
	}

	if curr.Val == target {
		*path = append(*path, curr.Val)
		return true
	}

	isPath := solve(curr.Left, target, path) || solve(curr.Right, target, path)
	if isPath {
		*path = append(*path, curr.Val)
	}

	return isPath
}

func reverse(s []int) {
	l := 0
	r := len(s) - 1

	for l < r {
		// Swap elements
		s[l], s[r] = s[r], s[l]

		// Move pointers toward the middle
		l++
		r--
	}
}

func main() {
	/*
	      10
	     /  \
	    5    20
	   / \     \
	  3   8     30
	*/
	root := &Node{
		Val: 10,
		Left: &Node{
			Val:   5,
			Left:  &Node{Val: 3},
			Right: &Node{Val: 8},
		},
		Right: &Node{
			Val:   20,
			Right: &Node{Val: 30},
		},
	}

	fmt.Println("Path to 8:", GetPath(root, 8))   // [10 5 8]
	fmt.Println("Path to 30:", GetPath(root, 30)) // [10 20 30]
	fmt.Println("Path to 99:", GetPath(root, 99)) // []
}
