package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

/**
 * mergeTwoLists merges two sorted linked lists in O(N+M) time.
 * We use a dummy node to avoid complex head-handling logic.
 */
func mergeTwoLists(A *ListNode, B *ListNode) *ListNode {
	// If one list is empty, return the other immediately
	if A == nil {
		return B
	}
	if B == nil {
		return A
	}

	// 1. Create a dummy node to act as the starting anchor
	dummy := &ListNode{value: 0}

	// 2. 'tail' will always point to the last node in our merged list
	tail := dummy

	// 3. Iterate as long as both lists have nodes
	for A != nil && B != nil {
		if A.value <= B.value {
			// Attach A's node and move A forward
			tail.next = A
			A = A.next
		} else {
			// Attach B's node and move B forward
			tail.next = B
			B = B.next
		}
		// Move the tail of our merged list forward
		tail = tail.next
	}

	// 4. Stitch the remaining nodes
	// If one list finishes, we just link the rest of the other list
	if A != nil {
		tail.next = A
	} else {
		tail.next = B
	}

	// 5. The actual merged list starts from dummy.next
	return dummy.next
}

func main() {
	// A = 5 -> 8 -> 20
	listA := &ListNode{5, &ListNode{8, &ListNode{20, nil}}}
	// B = 4 -> 11 -> 15
	listB := &ListNode{4, &ListNode{11, &ListNode{15, nil}}}

	merged := mergeTwoLists(listA, listB)

	// Print Result
	fmt.Print("Merged List: ")
	for merged != nil {
		fmt.Printf("%d", merged.value)
		if merged.next != nil {
			fmt.Print(" -> ")
		}
		merged = merged.next
	}
	fmt.Println()
}
