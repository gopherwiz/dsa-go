package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func rightSideView(root *Node) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		// Iterate through the current level
		for i := 0; i < levelSize; i++ {
			// Dequeue
			curr := queue[0]  // queue:peek
			queue = queue[1:] // queue:dequeue

			// If it's the LAST node in this level's count, add to result
			if i == levelSize-1 {
				result = append(result, curr.Value)
			}

			// Standard BFS: push children
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	
	return result
}

func main() {
	// Tree Structure:
	//      1  <-- View
	//     / \
	//    2   3 <-- View
	//     \   \
	//      5   4 <-- View
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Right = &Node{Value: 5}
	root.Right.Right = &Node{Value: 4}

	res := rightSideView(root)
	fmt.Printf("Right Side View: %v\n", res)
	// Output: [1 3 4]
}
