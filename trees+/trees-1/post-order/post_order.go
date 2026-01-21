package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func postorderRecursive(root *Node) []int {
	var result []int

	var traverse func(*Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}
		// 1. Visit Left
		traverse(node.Left)
		// 2. Visit Right
		traverse(node.Right)
		// 3. Visit Root (Current Node)
		result = append(result, node.Value)
	}

	traverse(root)

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

	arr := postorderRecursive(root)

	// Post-order Output should be: [4 5 2 3 1]
	fmt.Printf("Post-order Result: %v\n", arr)
}
