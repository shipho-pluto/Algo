package main

import "fmt"

type NodeTreeS struct {
	value *TreeS
	next  *NodeTreeS
}
type StackTree struct {
	top *NodeTreeS
}

func (s *StackTree) Add(value *TreeS) {
	newNode := &NodeTreeS{value: value, next: s.top}
	s.top = newNode
}
func (s *StackTree) Remove() *TreeS {
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

func DFS(t *TreeS, target int) bool {
	stack := StackTree{}
	stack.Add(t)

	for !stack.IsEmpty() {
		node := stack.Remove()
		if node.value == target {
			return true
		}

		if node.right != nil {
			stack.Add(node.right)
		}
		if node.left != nil {
			stack.Add(node.left)
		}
	}
	return false
}

type TreeS struct {
	value int
	left  *TreeS
	right *TreeS
}

func (t *TreeS) New(value int, left, right *TreeS) {
	t.value = value
	t.left = left
	t.right = right
}

func main() {
	tree := &TreeS{value: 20,
		left: &TreeS{value: 7,
			left: &TreeS{value: 4,
				left: nil,
				right: &TreeS{value: 6,
					left:  nil,
					right: nil,
				},
			},
			right: &TreeS{value: 9,
				left:  nil,
				right: nil,
			},
		},
		right: &TreeS{value: 35,
			left: &TreeS{value: 31,
				left: &TreeS{value: 28,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &TreeS{value: 40,
				left: &TreeS{value: 38,
					left:  nil,
					right: nil,
				},
				right: &TreeS{value: 52,
					left:  nil,
					right: nil,
				},
			},
		},
	}
	fmt.Println(DFS(tree, 52))
}
