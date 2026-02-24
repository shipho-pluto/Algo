package main

import (
	t "Algoritm/winterArc/trees/struct"
	"fmt"
)

func main() {
	tree := &t.TreeInt{Value: 0,
		Left: &t.TreeInt{Value: -8,
			Left: &t.TreeInt{Value: -10},
			Right: &t.TreeInt{Value: -6,
				Right: &t.TreeInt{Value: -6},
			},
		},
		Right: &t.TreeInt{Value: 100,
			Left: &t.TreeInt{Value: 8,
				Left:  &t.TreeInt{Value: 10},
				Right: &t.TreeInt{Value: 6},
			},
		},
	}
	fmt.Println(maxSumTree(tree))
}

func maxSumTree(tree *t.TreeInt) any {
	st := make(map[int]bool)
	_ = walkForSum(&st, tree)
	ans := 0
	for k := range st {
		ans = max(ans, k)
	}
	return ans
}

func walkForSum(m *map[int]bool, root *t.TreeInt) int {
	if root == nil {
		return 0
	}

	ans := root.Value
	l := walkForSum(m, root.Left)
	if l > 0 {
		ans += l
	}
	r := walkForSum(m, root.Right)
	if r > 0 {
		ans += r
	}

	(*m)[ans] = true
	return ans
}
