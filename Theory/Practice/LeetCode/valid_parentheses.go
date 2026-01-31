package main

import (
	"fmt"
)

type NodeValidParentheses struct {
	value interface{}
	next  *NodeValidParentheses
}

type StackValidParentheses struct {
	top *NodeValidParentheses
}

func (s *StackValidParentheses) Add(value interface{}) {
	newNode := &NodeValidParentheses{value: value, next: s.top}
	s.top = newNode
}

func (s *StackValidParentheses) Remove() interface{} {
	if s.top == nil {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

func (s *StackValidParentheses) IsEmpty() bool {
	return s.top == nil
}

func isValid(s string) bool {
	st := StackValidParentheses{}
	for _, x := range s {
		switch x {
		case '(':
			st.Add(x)
		case ')':
			if st.Remove() != '(' {
				return false
			}
		case '[':
			st.Add(x)
		case ']':
			if st.Remove() != '[' {
				return false
			}
		case '{':
			st.Add(x)
		case '}':
			if st.Remove() != '{' {
				return false
			}
		}
	}

	return st.IsEmpty()
}

func main() {
	arr := "(()){}[][][({])]"
	fmt.Println(isValid(arr))
}
