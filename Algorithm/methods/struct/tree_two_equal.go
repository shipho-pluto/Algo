package main

import (
	t "Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func main() {
	tree := &t.Tree{Value: 'a',
		Left: &t.Tree{Value: 'b',
			Left: &t.Tree{Value: 't'},
			Right: &t.Tree{Value: 'u',
				Right: &t.Tree{Value: 'e'},
			},
		},
		Right: &t.Tree{Value: 'a',
			Left: &t.Tree{Value: 'b',
				Left:  &t.Tree{Value: 'c'},
				Right: &t.Tree{Value: 'd'},
			},
		},
	}

	fmt.Println(returnTwoTree(tree))
}

func walkForTree(st *map[int][2]*t.Tree, root *t.Tree) int {
	if root == nil {
		return 0
	}

	ans := 1 << (root.Value - 'a')
	ans |= walkForTree(st, root.Left)
	ans |= walkForTree(st, root.Right)

	node := (*st)[ans]
	if node[0] == nil {
		node[0] = root
	} else if node[1] == nil {
		node[1] = root
	}
	(*st)[ans] = node

	return ans
}

func returnTwoTree(tree *t.Tree) [2]*t.Tree {
	var res [2]*t.Tree
	st := make(map[int][2]*t.Tree)
	_ = walkForTree(&st, tree)
	for _, node := range st {
		if node[0] != nil && node[1] != nil {
			return node
		}
	}

	return res
}
