package Structs

import (
	"strconv"
	"strings"
)

type Tree struct {
	Value int
	Right *Tree
	Left  *Tree
}

type TreeR struct {
	Value rune
	Right *TreeR
	Left  *TreeR
}

func (t *Tree) LEFT() int {
	val := t.Value
	if t.Left != nil {
		val = min(t.Left.LEFT(), val)
	}
	return val
}

func (t *Tree) RIGHT() int {
	val := t.Value
	if t.Right != nil {
		val = max(val, t.Right.RIGHT())
	}
	return val
}

type QueueTree struct {
	back  *Node
	front *Node
}
type Node struct {
	value *Tree
	next  *Node
}

func (q *QueueTree) Add(value *Tree) {
	newNode := &Node{value: value}
	if q.back == nil {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}
}
func (q *QueueTree) Remove() *Tree {
	if q.front == nil {
		return nil
	}
	value := q.front.value
	q.front = q.front.next
	if q.front == nil {
		q.back = nil
	}
	return value
}
func (s *QueueTree) IsEmpty() bool {
	return s.back == nil
}

func (t *Tree) Print() string {
	queue := QueueTree{}
	queue.Add(t)
	res := strings.Builder{}

	for !queue.IsEmpty() {
		node := queue.Remove()
		res.WriteString(strconv.Itoa(node.Value) + " ")

		if node.Left != nil {
			queue.Add(node.Left)
		}
		if node.Right != nil {
			queue.Add(node.Right)
		}
	}
	return res.String()
}
