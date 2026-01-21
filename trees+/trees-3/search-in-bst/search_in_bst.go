package main

// searchRecursive uses the call stack to find the target.
func searchRecursive(root *Node, target int) *Node {
	// Base case: nil or found
	if root == nil || root.Val == target {
		return root
	}

	// If target is smaller, go left
	if target < root.Val {
		return searchRecursive(root.Left, target)
	}

	// If target is larger, go right
	return searchRecursive(root.Right, target)
}

//func main() {
//	/*
//	   Constructing this BST:
//	         10
//	       /    \
//	      5      15
//	     / \    /  \
//	    2   7  12   20
//	*/
//	root := &Node{Val: 10}
//	root.Left = &Node{Val: 5, Left: &Node{Val: 2}, Right: &Node{Val: 7}}
//	root.Right = &Node{Val: 15, Left: &Node{Val: 12}, Right: &Node{Val: 20}}
//
//	targets := []int{7, 12, 100}
//
//	fmt.Println("Searching BST:")
//	for _, t := range targets {
//		// Using the iterative search
//		result := searchRecursive(root, t)
//
//		if result != nil {
//			fmt.Printf("Target %d: FOUND at node address %p\n", t, result)
//		} else {
//			fmt.Printf("Target %d: NOT FOUND\n", t)
//		}
//	}
//}
