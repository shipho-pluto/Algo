package main

import (
	t "Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func main() {
	tree := &t.Tree{Value: 0,
		Left: &t.Tree{Value: -8,
			Left: &t.Tree{Value: -10},
			Right: &t.Tree{Value: -6,
				Right: &t.Tree{Value: -6},
			},
		},
		Right: &t.Tree{Value: 100,
			Left: &t.Tree{Value: 8,
				Left:  &t.Tree{Value: 10},
				Right: &t.Tree{Value: 6},
			},
		},
	}
	fmt.Println(maxSumTree(tree))
}

func maxSumTree(root *t.Tree) int {
	st := make(map[int]bool)
	_ = walkForSumTree(root, &st)
	res := root.Value
	for sum := range st {
		res = max(res, sum)
	}
	return res
}

func walkForSumTree(root *t.Tree, m *map[int]bool) int {
	if root == nil {
		return 0
	}

	ans := root.Value
	ansL := walkForSumTree(root.Left, m)
	ansR := walkForSumTree(root.Right, m)
	if ansL > 0 {
		ans += ansL
	}
	if ansR > 0 {
		ans += ansR
	}

	(*m)[ans] = true

	return ans
}
