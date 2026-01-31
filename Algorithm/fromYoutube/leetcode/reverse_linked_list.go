package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(reverseList(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	top *ListNode
}

func reverseList(head *ListNode) *ListNode {
	st := Stack{}
	for ; head != nil; head = head.Next {
		node := &ListNode{head.Val, st.top}
		st.top = node
	}
	return st.top
}
