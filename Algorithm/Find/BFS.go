package main

import "fmt"

type QueueTree struct {
	back  *NodeTreeQ
	front *NodeTreeQ
}
type NodeTreeQ struct {
	value *TreeQ
	next  *NodeTreeQ
}

func (q *QueueTree) Add(value *TreeQ) {
	newNode := &NodeTreeQ{value: value}
	if q.back == nil {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}
}
func (q *QueueTree) Remove() *TreeQ {
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

func BFS(t *TreeQ, target int) bool {
	queue := QueueTree{}
	queue.Add(t)

	for !queue.IsEmpty() {
		node := queue.Remove()
		if node.value == target {
			return true
		}

		if node.right != nil {
			queue.Add(node.right)
		}
		if node.left != nil {
			queue.Add(node.left)
		}
	}
	return false
}

type TreeQ struct {
	value int
	left  *TreeQ
	right *TreeQ
}

func (t *TreeQ) New(value int, left, right *TreeQ) {
	t.value = value
	t.left = left
	t.right = right
}

func main() {
	tree := &TreeQ{value: 20,
		left: &TreeQ{value: 7,
			left: &TreeQ{value: 4,
				left: nil,
				right: &TreeQ{value: 6,
					left:  nil,
					right: nil,
				},
			},
			right: &TreeQ{value: 9,
				left:  nil,
				right: nil,
			},
		},
		right: &TreeQ{value: 35,
			left: &TreeQ{value: 31,
				left: &TreeQ{value: 28,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &TreeQ{value: 40,
				left: &TreeQ{value: 38,
					left:  nil,
					right: nil,
				},
				right: &TreeQ{value: 52,
					left:  nil,
					right: nil,
				},
			},
		},
	}
	fmt.Println(BFS(tree, 52))
}
