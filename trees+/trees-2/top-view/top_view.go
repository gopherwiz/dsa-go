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

func topView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// hm will store only the FIRST node seen at each level
	hm := make(map[int]int)
	queue := []Pair{{node: root, level: 0}}

	minLevel, maxLevel := 0, 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		node := curr.node
		level := curr.level

		// Update boundaries
		if level < minLevel {
			minLevel = level
		}
		if level > maxLevel {
			maxLevel = level
		}

		// ONLY add to map if the level has never been seen before
		if _, exists := hm[level]; !exists {
			hm[level] = node.Val
		}

		if node.Left != nil {
			queue = append(queue, Pair{node: node.Left, level: level - 1})
		}
		if node.Right != nil {
			queue = append(queue, Pair{node: node.Right, level: level + 1})
		}
	}

	// Build the final result slice
	result := make([]int, 0, maxLevel-minLevel+1)
	for i := minLevel; i <= maxLevel; i++ {
		result = append(result, hm[i])
	}

	return result
}

func main() {
	// Tree:
	//      1
	//    /   \
	//   2     3
	//    \   /
	//     5 6
	// The Top View should be [2, 1, 3]
	// Note: 5 and 6 are hidden by the nodes above them.
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}
	root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}

	fmt.Printf("Top View: %v\n", topView(root))
}
