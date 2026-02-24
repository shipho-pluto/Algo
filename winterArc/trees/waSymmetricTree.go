package main

import (
	t "Algoritm/winterArc/trees/struct"
	"fmt"
)

func main() {
	tree := &t.TreeInt{Value: 1,
		Right: &t.TreeInt{Value: 2,
			Right: &t.TreeInt{Value: 3,
				Right: &t.TreeInt{Value: 10},
				Left:  &t.TreeInt{Value: 9},
			},
			Left: &t.TreeInt{Value: 4,
				Right: &t.TreeInt{Value: 8},
				Left:  &t.TreeInt{Value: 7},
			},
		},
		Left: &t.TreeInt{Value: 2,
			Right: &t.TreeInt{Value: 4,
				Right: &t.TreeInt{Value: 7},
				Left:  &t.TreeInt{Value: 8},
			},
			Left: &t.TreeInt{Value: 3,
				Right: &t.TreeInt{Value: 9},
				Left:  &t.TreeInt{Value: 10},
			},
		},
	}

	fmt.Println(isTreeSymmetric(tree))
}

type q1 []*t.TreeInt

func pow(i int) int {
	res := 1
	p := 2
	for range i {
		res *= p
	}
	return res
}

func isTreeSymmetric(root *t.TreeInt) bool {
	queue := make(q1, 0)
	queue = append(queue, root)
	arr := make([]int, 0, 2)
	step := 1
	for i := 1; len(queue) != 0; i++ {
		root = queue[0]
		queue = queue[1:]

		if root.Left != nil {
			queue = append(queue, root.Left)
			arr = append(arr, root.Left.Value)
		}

		if root.Right != nil {
			queue = append(queue, root.Right)
			arr = append(arr, root.Right.Value)
		}

		if len(arr) == pow(step) {
			for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
				if arr[l] != arr[r] {
					return false
				}
			}
			step += 1
			arr = make([]int, 0, pow(step))
		}
	}
	return true
}

func isTreeSymmetricEasy(tree *t.TreeInt) any {
	return walk(tree.Left, tree.Right)
}

func walk(l *t.TreeInt, r *t.TreeInt) bool {
	if l != nil && r != nil {
		if l.Value != r.Value {
			return false
		}
		return walk(l.Left, r.Right) && walk(l.Right, r.Left)
	} else if l == nil && r == nil {
		return true
	}
	return false
}
