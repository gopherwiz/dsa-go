package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func preorderRecursive(root *Node) []int {
	var result []int

	var traverse func(*Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}
		// 1. Visit Root (Current Node) First
		result = append(result, node.Value)

		// 2. Visit Left
		traverse(node.Left)

		// 3. Visit Right
		traverse(node.Right)
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

	arr := preorderRecursive(root)

	// Pre-order Output: [1 2 4 5 3]
	fmt.Printf("Pre-order Result: %v\n", arr)
}
