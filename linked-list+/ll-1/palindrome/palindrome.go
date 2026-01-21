package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// isPalindrome checks if a linked list is a palindrome in O(N) time and O(1) space.
func isPalindrome(head *ListNode) bool {
	// Base case: empty or single node list is always a palindrome
	if head == nil || head.Next == nil {
		return true
	}

	// 1. Find the middle of the list (Middle 1 for even lengths)
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. Reverse the second half
	// slow.Next is the start of the second half
	newHead := reverseList(slow.Next)

	// 3. Compare the two halves
	first := head
	second := newHead
	result := true
	for second != nil {
		if first.Val != second.Val {
			result = false
			break
		}
		first = first.Next
		second = second.Next
	}

	// 4. Restore the list (Good practice in interviews)
	reverseList(newHead)

	return result
}

// reverseList is a helper to reverse a linked list in-place
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

func main() {
	// Example: 1 -> 2 -> 3 -> 2 -> 1
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 2}
	n5 := &ListNode{Val: 1}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	fmt.Println("Is Palindrome:", isPalindrome(n1)) // Expected: true
}
