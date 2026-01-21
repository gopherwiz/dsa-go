package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func inorderRecursive(root *Node) []int {
	// Base case
	if root == nil {
		return nil
	}

	// Step 1: Get results from Left
	result := inorderRecursive(root.Left)

	// Step 2: Directly append Current Value
	result = append(result, root.Value)

	// Step 3: Directly append results from Right
	// The "..." expands the slice into individual arguments
	result = append(result, inorderRecursive(root.Right)...)

	return result
}

func main() {
	root := &Node{Value: 1,
		Left:  &Node{Value: 2, Left: &Node{Value: 4}, Right: &Node{Value: 5}},
		Right: &Node{Value: 3},
	}

	arr := inorderRecursive(root)

	fmt.Printf("Recursive Result: %v\n", arr)
	// Output: [4 2 5 1 3]
}
