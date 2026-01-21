package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	node  *TreeNode
	level int
}

func verticalOrderTraversal(A *TreeNode) [][]int {
	if A == nil {
		return [][]int{}
	}

	queue := []Pair{{node: A, level: 0}}
	hm := make(map[int][]int)

	// Track range to avoid sorting keys later
	minLevel := 0
	maxLevel := 0

	for len(queue) > 0 {
		// Pop front
		curr := queue[0]
		queue = queue[1:]

		node := curr.node
		level := curr.level

		// Update min/max
		if level < minLevel {
			minLevel = level
		}
		if level > maxLevel {
			maxLevel = level
		}

		// Add to map
		hm[level] = append(hm[level], node.Val)

		// Push children
		if node.Left != nil {
			queue = append(queue, Pair{node: node.Left, level: level - 1})
		}
		if node.Right != nil {
			queue = append(queue, Pair{node: node.Right, level: level + 1})
		}
	}

	// Build result slice from min to max
	result := make([][]int, 0, maxLevel-minLevel+1)
	for i := minLevel; i <= maxLevel; i++ {
		result = append(result, hm[i])
	}

	return result
}

func main() {
	// Example 1 construction
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}}
	root.Right = &TreeNode{Val: 7, Right: &TreeNode{Val: 9}}

	fmt.Println(verticalOrderTraversal(root))
	// Output: [[2] [3] [6 5] [7] [9]]
}
