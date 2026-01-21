package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func reverseBetween(A *ListNode, B int, C int) *ListNode {
	if A == nil || B == C {
		return A
	}

	// Dummy node handles the case where B = 1
	dummy := &ListNode{next: A}
	beforeStart := dummy

	// Step 1: Walk to the node just BEFORE the reversal starts
	for i := 1; i < B; i++ {
		beforeStart = beforeStart.next
	}

	// Step 2: Set up pointers for reversal
	// 'start' will eventually be the tail of our reversed segment
	start := beforeStart.next
	var prev *ListNode = nil
	curr := start

	// Step 3: Standard one-pass reversal for (C - B + 1) nodes
	for i := 0; i < (C - B + 1); i++ {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	// Step 4: Stitch the reversed segment back into the main list
	// beforeStart.next (node B-1) now points to the new head (prev)
	beforeStart.next = prev
	// start.next (the original node B) now points to the remaining list (curr)
	start.next = curr

	return dummy.next
}

func main() {
	// Creating list manually: 1 -> 2 -> 3 -> 4 -> 5
	node5 := &ListNode{value: 5, next: nil}
	node4 := &ListNode{value: 4, next: node5}
	node3 := &ListNode{value: 3, next: node4}
	node2 := &ListNode{value: 2, next: node3}
	head := &ListNode{value: 1, next: node2}

	B, C := 2, 4
	fmt.Printf("Reversing from position %d to %d...\n", B, C)

	// Execute reversal
	newHead := reverseBetween(head, B, C)

	// Print directly in main
	fmt.Print("Result: ")
	temp := newHead
	for temp != nil {
		fmt.Printf("%d", temp.value)
		if temp.next != nil {
			fmt.Print(" -> ")
		}
		temp = temp.next
	}
	fmt.Println()
}
