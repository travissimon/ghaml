package main

import ()

type stack struct {
	top  *StackNode
	size int
}

type StackNode struct {
	value interface{}
	next  *StackNode
}

// Return the stack's length
func (s *stack) count() int {
	return s.size
}

// Push a new element onto the stack
func (s *stack) push(value interface{}) {
	s.top = &StackNode{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *stack) pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

// combines a pop and a push to preview an item
func (s *stack) peek() (value interface{}) {
	item := s.pop()
	if item != nil {
		s.push(item)
	}
	return item
}
