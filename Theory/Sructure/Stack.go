package Struct

type nodeS struct {
	value interface{}
	next  *nodeS
}

type Stack struct {
	top *nodeS
}

func (s *Stack) Push(value interface{}) {
	newNode := &nodeS{value: value, next: s.top}
	s.top = newNode
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

func (s *Stack) Look() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.value
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
