package main

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Add(value interface{}) {
	newNode := &Node{value: value, next: s.top}
	s.top = newNode
}

func (s *Stack) Remove() interface{} {
	if s.top == nil {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
