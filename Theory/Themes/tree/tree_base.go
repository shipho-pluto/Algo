package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) FindElement(x int) int {
	if n == nil {
		return -1
	}
	index := 0
	if x == n.Val {
		return index
	} else if x < n.Val {
		index = n.Left.FindElement(x)
	} else {
		index = n.Right.FindElement(x)
	}

	return index
}

func (n *Node) AddElement(x int, head *Node) {
	if n == nil {
		if x < head.Val {
			head.Left = &Node{Val: x}
		} else {
			head.Right = &Node{Val: x}
		}
	} else {
		if x < n.Val {
			n.Left.AddElement(x, n)
		} else {
			n.Right.AddElement(x, n)
		}
	}
}

func (n *Node) RemoveElementNoSon(x int, head *Node) {
	if x == n.Val {
		if x < head.Val {
			head.Left = nil
		} else {
			head.Right = nil
		}
	} else if x < n.Val {
		n.Left.RemoveElementNoSon(x, n)
	} else {
		n.Right.RemoveElementNoSon(x, n)
	}
}

func (n *Node) RemoveElementOneSon(x int, head *Node) {
	if x == n.Val {
		if x < head.Val {
			if n.Right == nil {
				head.Left = n.Left
			} else {
				head.Left = n.Right
			}
		} else {
			if n.Right == nil {
				head.Right = n.Left
			} else {
				head.Right = n.Right
			}
		}
	} else if x < n.Val {
		n.Left.RemoveElementOneSon(x, n)
	} else {
		n.Right.RemoveElementOneSon(x, n)
	}
}

func (n *Node) goLeft(head *Node) int {
	if n.Left != nil {
		n.Left.goLeft(n)
	}
	if n.Right != nil {
		head.Left = n.Right
	}
	return n.Val
}

func (n *Node) goRight(head *Node) int {
	if n.Right != nil {
		n.Right.goRight(n)
	}
	if n.Left != nil {
		head.Right = n.Left
	}
	return n.Val
}

func (n *Node) RemoveElementTwoSon(x int, head *Node) {
	if x == n.Val {
		if x < head.Val {
			head.Left = &Node{Val: n.Right.goLeft(n), Left: n.Left, Right: n.Right}
		} else {
			head.Right = &Node{Val: n.Right.goLeft(n), Left: n.Left, Right: n.Right}
		}
	} else if x < n.Val {
		n.Left.RemoveElementTwoSon(x, n)
	} else {
		n.Right.RemoveElementTwoSon(x, n)
	}
}

func (n *Node) FindRecursive() {
	if n.Left != nil {
		n.Left.FindRecursive()
	}
	if n.Right != nil {
		n.Right.FindRecursive()
	}
	fmt.Printf("%d ", n.Val)
}
