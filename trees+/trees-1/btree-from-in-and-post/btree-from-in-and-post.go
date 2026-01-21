package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Build function: Reconstructs tree from Inorder and Postorder
func build(in []int, post []int) *Node {
	if len(in) == 0 {
		return nil
	}

	// 1. Pick the Root (Last of postorder)
	val := post[len(post)-1]
	root := &Node{Value: val}

	// 2. Find where it splits the Inorder list
	i := 0
	for i < len(in) && in[i] != val {
		i++
	}

	// 3. Slice and Recurse
	root.Left = build(in[:i], post[:i])
	root.Right = build(in[i+1:], post[i:len(post)-1])

	return root
}

// Verification function: Returns Inorder slice from tree
func getInorder(root *Node) []int {
	var res []int
	if root == nil {
		return nil
	}

	res = append(res, getInorder(root.Left)...)
	res = append(res, root.Value)
	res = append(res, getInorder(root.Right)...)

	return res
}

func main() {
	// Original data used for building
	originalInorder := []int{4, 2, 5, 1, 3}
	postorder := []int{4, 5, 2, 3, 1}

	// Step 1: Build the tree
	root := build(originalInorder, postorder)

	// Step 2: Verify by traversing the NEW tree
	verificationInorder := getInorder(root)

	fmt.Printf("Original Inorder:     %v\n", originalInorder)
	fmt.Printf("Verification Inorder: %v\n", verificationInorder)

	// Final Check
	if reflect.DeepEqual(originalInorder, verificationInorder) {
		fmt.Println("✅ Verification Successful: The tree was built correctly!")
	} else {
		fmt.Println("❌ Verification Failed: The tree structure is incorrect.")
	}
}
