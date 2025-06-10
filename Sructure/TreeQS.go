package main

import "fmt"

type QueueTree struct {
	back  *NodeTreeQ
	front *NodeTreeQ
}
type NodeTreeQ struct {
	value *TreeTwo
	next  *NodeTreeQ
}

func (q *QueueTree) Add(value *TreeTwo) {
	newNode := &NodeTreeQ{value: value}
	if q.back == nil {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}
}
func (q *QueueTree) Remove() *TreeTwo {
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

type NodeTreeS struct {
	value *TreeTwo
	next  *NodeTreeS
}
type StackTree struct {
	top *NodeTreeS
}

func (s *StackTree) Add(value *TreeTwo) {
	newNode := &NodeTreeS{value: value, next: s.top}
	s.top = newNode
}
func (s *StackTree) Remove() *TreeTwo {
	if s.top == nil {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	return value
}
func (s *StackTree) IsEmpty() bool {
	return s.top == nil
}

func (t *TreeTwo) DFS() int {
	stack := StackTree{}
	stack.Add(t)

	sum := 0

	for !stack.IsEmpty() {
		node := stack.Remove()
		sum += node.value

		if node.right != nil {
			stack.Add(node.right)
		}
		if node.left != nil {
			stack.Add(node.left)
		}
	}
	return sum
}

func (t *TreeTwo) BFS() int {
	queue := QueueTree{}
	queue.Add(t)

	sum := 0

	for !queue.IsEmpty() {
		node := queue.Remove()
		sum += node.value

		if node.right != nil {
			queue.Add(node.right)
		}
		if node.left != nil {
			queue.Add(node.left)
		}
	}
	return sum
}

type TreeTwo struct {
	value int
	left  *TreeTwo
	right *TreeTwo
}

func (t *TreeTwo) New(value int, left, right *TreeTwo) {
	t.value = value
	t.left = left
	t.right = right
}

func main() {
	tree := &TreeTwo{value: 20,
		left: &TreeTwo{value: 7,
			left: &TreeTwo{value: 4,
				left: nil,
				right: &TreeTwo{value: 6,
					left:  nil,
					right: nil,
				},
			},
			right: &TreeTwo{value: 9,
				left:  nil,
				right: nil,
			},
		},
		right: &TreeTwo{value: 35,
			left: &TreeTwo{value: 31,
				left: &TreeTwo{value: 28,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &TreeTwo{value: 40,
				left: &TreeTwo{value: 38,
					left:  nil,
					right: nil,
				},
				right: &TreeTwo{value: 52,
					left:  nil,
					right: nil,
				},
			},
		},
	}
	fmt.Println(tree.DFS())
	fmt.Println(tree.BFS())
}
