package main

func rightSideViewDFS(root *Node) []int {
	var result []int
	// Start recursion at level 0
	dfs(root, 0, &result)
	return result
}

func dfs(node *Node, level int, result *[]int) {
	if node == nil {
		return
	}

	// Logic: If the current level equals the number of elements in result,
	// it means we are visiting this depth for the first time.
	// Since we visit Right before Left, this MUST be the rightmost node.
	if level == len(*result) {
		*result = append(*result, node.Value)
	}

	// Visit RIGHT child first
	dfs(node.Right, level+1, result)
	// Visit LEFT child second
	dfs(node.Left, level+1, result)
}

//func main() {
//	// Tree:
//	//      1
//	//     / \
//	//    2   3
//	//     \   \
//	//      5   4
//	root := &Node{Value: 1}
//	root.Left = &Node{Value: 2}
//	root.Right = &Node{Value: 3}
//	root.Left.Right = &Node{Value: 5}
//	root.Right.Right = &Node{Value: 4}
//
//	fmt.Printf("Right View (DFS): %v\n", rightSideView(root))
//}
