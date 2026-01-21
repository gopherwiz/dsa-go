package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * detectCycle finds the starting node of a cycle in a linked list.
 * Time Complexity: O(N)
 * Space Complexity: O(1)
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	// Step 1: Detect if a loop exists using Tortoise and Hare
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// If they meet, a loop is confirmed
		// Step 2: Find the starting point of the loop
		if slow == fast {
			// Reset slow to head, keep fast at the collision point
			slow = head
			// Move both one step at a time until they meet again
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}

			// The node where they meet again is the starting node of the cycle
			return slow
		}
	}

	// If no collision point was found, there's no cycle
	return nil
}

func main() {
	// Creating a list: 1 -> 2 -> 3 -> 4 -> 5 -> 3 (cycle starts at 3)
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n3 // Cycle back to n3

	startNode := detectCycle(n1)
	if startNode != nil {
		fmt.Printf("Cycle starts at node with value: %d\n", startNode.Val)
	} else {
		fmt.Println("No cycle found.")
	}
}
