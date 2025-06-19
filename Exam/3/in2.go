package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ll *ListNode) Print() string {
	str := ""
	if ll != nil {
		str += fmt.Sprint(ll.Val) + " "
		str += ll.Next.Print()
	}
	return str
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
	}
	if head == nil {
		return head
	}
	lNext := head
	rNext := head.Next
	for ; rNext != nil; rNext = rNext.Next {
		if rNext.Val == val {
			lNext.Next = rNext.Next
		} else {
			lNext.Next = rNext
			lNext = lNext.Next
		}
	}
	return head
}

func main() {
	head := &ListNode{6, &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}}
	fmt.Println(removeElements(head, 6))
}
