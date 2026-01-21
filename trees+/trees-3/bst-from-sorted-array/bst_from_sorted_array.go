package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// sortedArrayToBST is the main recursive function
func sortedArrayToBST(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	// 1. Find the middle index
	mid := len(nums) / 2

	// 2. Create the root node with the middle element
	root := &Node{Val: nums[mid]}

	// 3. Recursively build the left and right subtrees
	// Left half: all elements before mid
	root.Left = sortedArrayToBST(nums[:mid])
	// Right half: all elements after mid
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

// Helper: Inorder traversal to verify the BST property
func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Printf("%d ", root.Val)
	inorder(root.Right)
}

func main() {
	// Sorted array with unique elements
	nums := []int{10, 20, 30, 40, 50, 60, 70}

	fmt.Printf("Original Array: %v\n", nums)

	root := sortedArrayToBST(nums)

	fmt.Print("Inorder Traversal of Resulting BST: ")
	inorder(root)
	fmt.Println("\n(If this matches the original array, the BST is valid!)")
}
