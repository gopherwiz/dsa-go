package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

/**
 * hasCycle returns true if the linked list contains a loop.
 * Time Complexity: O(N)
 * Space Complexity: O(1)
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow := head
	fast := head

	// Move fast by 2 and slow by 1
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		// If they meet, a cycle exists
		if slow == fast {
			return true
		}
	}

	// If fast reaches the end, no cycle
	return false
}

func main() {
	// Setup: 1 -> 2 -> 3 -> 4 (back to 2)
	n1 := &ListNode{value: 1}
	n2 := &ListNode{value: 2}
	n3 := &ListNode{value: 3}
	n4 := &ListNode{value: 4}

	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n2 // Create cycle at node 2

	if hasCycle(n1) {
		fmt.Println("Cycle detected!")
	} else {
		fmt.Println("No cycle detected.")
	}
}
