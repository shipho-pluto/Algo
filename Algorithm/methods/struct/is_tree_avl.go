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
		Right: &t.Tree{Value: 0,
			Left: &t.Tree{Value: -8,
				Left:  &t.Tree{Value: -10},
				Right: &t.Tree{Value: -6},
			},
		},
	}
	fmt.Println(isTreeAVL(tree))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isTreeAVL(tree *t.Tree) bool {
	if tree == nil {
		return false
	}
	return abs(high(tree.Left)-high(tree.Right)) <= 1
}

func high(root *t.Tree) int {
	if root == nil {
		return 0
	}
	return max(high(root.Left), high(root.Right)) + 1
}
