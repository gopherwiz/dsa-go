package main

import "fmt"

type UserStack struct {
	queue []int
}

func NewUserStack() *UserStack {
	return &UserStack{
		queue: make([]int, 0),
	}
}

// Push element x onto stack
func (s *UserStack) Push(x int) {
	// 1. Record the size before adding the new element
	size := len(s.queue)

	// 2. Add the new element to the back
	s.queue = append(s.queue, x)

	// 3. Rotate the queue: move all previous elements to the back
	// This puts the newly added element at the front (index 0)
	for i := 0; i < size; i++ {
		// Get front element
		front := s.queue[0]
		// Remove front
		s.queue = s.queue[1:]
		// Push to back
		s.queue = append(s.queue, front)
	}
}

// Pop removes the element on top of the stack and returns it
func (s *UserStack) Pop() int {
	if s.Empty() {
		return -1
	}

	val := s.queue[0]
	s.queue = s.queue[1:]

	return val
}

// Top returns the top element
func (s *UserStack) Top() int {
	if s.Empty() {
		return -1
	}

	return s.queue[0]
}

func (s *UserStack) Empty() bool {
	return len(s.queue) == 0
}

func main() {
	stack := NewUserStack()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Top (should be 30):", stack.Top())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Top now (should be 20):", stack.Top())
}
