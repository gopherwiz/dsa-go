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

func bottomView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// Map stores the LATEST node seen at each level
	hm := make(map[int]int)
	queue := []Pair{{node: root, level: 0}}

	minLevel, maxLevel := 0, 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		node := curr.node
		level := curr.level

		if level < minLevel {
			minLevel = level
		}
		if level > maxLevel {
			maxLevel = level
		}

		// OVERWRITE the value for this level.
		// The last node processed at this level will be the bottom-most.
		hm[level] = node.Val

		if node.Left != nil {
			queue = append(queue, Pair{node: node.Left, level: level - 1})
		}
		if node.Right != nil {
			queue = append(queue, Pair{node: node.Right, level: level + 1})
		}
	}

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
	// Bottom View: [2, 5, 6, 3] or [2, 6, 5, 3] depending on 5 vs 6
	// In BFS, if 5 and 6 are at the same level, 6 will likely overwrite 5.
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}
	root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}

	fmt.Printf("Bottom View: %v\n", bottomView(root))
}
