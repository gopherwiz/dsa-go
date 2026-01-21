package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// levelOrder - here and all related questions, pref is given to left side first in traversal
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*Node{root} // Initialize queue with the root
	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []int
		// Process all nodes at the current level
		for i := 0; i < levelSize; i++ {
			// Dequeue: Get the first element
			node := queue[0]  // queue:peek
			queue = queue[1:] // queue:dequeue

			currentLevel = append(currentLevel, node.Value)

			// Enqueue children for the next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Add the finished level to our results
		result = append(result, currentLevel)
	}

	return result
}

func main() {
	// Tree Structure:
	//      1
	//     / \
	//    2   3
	//   / \
	//  4   5
	root := &Node{Value: 1,
		Left:  &Node{Value: 2, Left: &Node{Value: 4}, Right: &Node{Value: 5}},
		Right: &Node{Value: 3},
	}

	levels := levelOrder(root)

	fmt.Println("Level Order Traversal:")
	for i, level := range levels {
		fmt.Printf("Level %d: %v\n", i, level)
	}
}
