package main

import "fmt"

// UserQueue struct for queue implementation
type UserQueue struct {
	s1 []int // Input stack
	s2 []int // Output stack
}

// NewUserQueue initializes your data structure
func NewUserQueue() *UserQueue {
	return &UserQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

// Push element X to the back of queue
func (q *UserQueue) Push(x int) {
	// Always push to the input stack
	q.s1 = append(q.s1, x)
}

// Pop removes the element from in front of queue and returns that element
func (q *UserQueue) Pop() int {
	q.shiftStacks()

	if len(q.s2) == 0 {
		return -1 // Or handle error
	}

	// Pop from s2
	val := q.s2[len(q.s2)-1]
	q.s2 = q.s2[:len(q.s2)-1]

	return val
}

// Peek gets the front element of the queue
func (q *UserQueue) Peek() int {
	q.shiftStacks()

	if len(q.s2) == 0 {
		return -1
	}

	return q.s2[len(q.s2)-1]
}

// Empty returns whether the queue is empty
func (q *UserQueue) Empty() bool {
	return len(q.s1) == 0 && len(q.s2) == 0
}

// shiftStacks moves elements from s1 to s2 only if s2 is empty
func (q *UserQueue) shiftStacks() {
	if len(q.s2) == 0 {
		for len(q.s1) > 0 {
			// Pop from s1
			val := q.s1[len(q.s1)-1]
			q.s1 = q.s1[:len(q.s1)-1]

			// Push to s2
			q.s2 = append(q.s2, val)
		}
	}
}

func main() {
	q := NewUserQueue()
	q.Push(1)
	q.Push(2)
	fmt.Println("Peek:", q.Peek())   // Output: 1
	fmt.Println("Pop:", q.Pop())     // Output: 1
	fmt.Println("Empty:", q.Empty()) // Output: false
}
