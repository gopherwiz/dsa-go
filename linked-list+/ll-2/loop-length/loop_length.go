package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * countNodesInLoop returns the length of the loop in a linked list.
 * If no loop exists, it returns 0.
 */
func countNodesInLoop(head *ListNode) int {
	if head == nil || head.Next == nil {
		return 0
	}

	slow := head
	fast := head

	// Step 1: Detect if a loop exists
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// If they meet, a loop is found
		if slow == fast {
			return findLength(slow)
		}
	}

	// No loop found
	return 0
}

/**
 * findLength calculates the number of nodes in the cycle.
 */
func findLength(slow *ListNode) int {
	count := 1
	curr := slow.Next

	// Move one step at a time until reaching the start node again
	for curr != slow {
		count++
		curr = curr.Next
	}

	return count
}

func main() {
	// Creating a list: 1 -> 2 -> 3 -> 4 -> 5 -> 3 (loop length is 3)
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n3 // Cycle starts back at n3

	fmt.Printf("Length of the loop: %d\n", countNodesInLoop(n1))
}
