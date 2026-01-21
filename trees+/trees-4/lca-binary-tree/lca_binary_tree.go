package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func getPath(root *Node, target int) []int {
	var path []int
	commonPath(root, target, &path)

	reverse(path)

	return path
}

func commonPath(curr *Node, target int, path *[]int) bool {
	if curr == nil {
		return false
	}

	if curr.Val == target {
		*path = append(*path, curr.Val)
		return true
	}

	isPath := commonPath(curr.Left, target, path) || commonPath(curr.Right, target, path)
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

func FindLCA(root *Node, p int, q int) int {
	path1 := getPath(root, p)
	path2 := getPath(root, q)

	// If either path is empty, one of the nodes doesn't exist
	if len(path1) == 0 || len(path2) == 0 {
		return -1
	}

	i := 0
	// Move i forward as long as both paths are identical
	for i < len(path1) && i < len(path2) && path1[i] == path2[i] {
		i++
	}

	// The loop stops at the first divergence.
	// The LCA is the element just before the divergence.
	return path1[i-1]
}

func main() {
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

	// LCA of 3 and 8 is 5
	fmt.Println("LCA of 3 and 8:", FindLCA(root, 3, 8))

	// LCA of 3 and 30 is 10
	fmt.Println("LCA of 3 and 30:", FindLCA(root, 3, 30))
}
