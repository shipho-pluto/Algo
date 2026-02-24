package main

import (
	t "Algoritm/winterArc/trees/struct"
	"fmt"
)

func main() {
	tree := &t.TreeCh{Value: 'a',
		Left: &t.TreeCh{Value: 'b',
			Left: &t.TreeCh{Value: 'c'},
			Right: &t.TreeCh{Value: 'd',
				Right: &t.TreeCh{Value: 'd'},
			},
		},
		Right: &t.TreeCh{Value: 'a',
			Left: &t.TreeCh{Value: 'b',
				Left:  &t.TreeCh{Value: 'c'},
				Right: &t.TreeCh{Value: 'd'},
			},
		},
	}

	fmt.Println(twoEqTrees(tree))
}

func twoEqTrees(tree *t.TreeCh) (*t.TreeCh, *t.TreeCh, int) {
	st := make(map[int][2]*t.TreeCh)
	_ = walkForTree(&st, tree)
	res := [2]*t.TreeCh{}
	maxH := 0
	for _, v := range st {
		h1, h2 := highTree(v[0]), highTree(v[1])
		if v[0] != nil && v[1] != nil {
			if maxH < h1+h2 {
				maxH = h1 + h2
				res = v
			}
		}
	}
	return res[0], res[1], maxH
}

func highTree(t *t.TreeCh) int {
	if t != nil {
		return max(highTree(t.Left)+1, highTree(t.Right)+1)
	}
	return 1
}

func walkForTree(m *map[int][2]*t.TreeCh, tree *t.TreeCh) int {
	if tree == nil {
		return 0
	}
	ans := 1 << int(tree.Value-'a')
	ans |= walkForTree(m, tree.Left)
	ans |= walkForTree(m, tree.Right)

	node := (*m)[ans]
	if node[0] == nil {
		node[0] = tree
	} else if node[1] == nil {
		node[1] = tree
	}
	(*m)[ans] = node
	return ans
}
