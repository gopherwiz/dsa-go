package main

func inorderIterative(root *Node) []int {
	var res []int
	var stack []*Node
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1] // stack:peek
		res = append(res, curr.Value)

		stack = stack[:len(stack)-1] // stack:pop

		curr = curr.Right
	}

	return res
}

//func main() {
//	root := &Node{Value: 1,
//		Left:  &Node{Value: 2, Left: &Node{Value: 4}, Right: &Node{Value: 5}},
//		Right: &Node{Value: 3},
//	}
//
//	fmt.Println("Iterative:", inorderIterative(root))
//}
