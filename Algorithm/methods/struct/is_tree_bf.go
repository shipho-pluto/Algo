package main

import (
	t "Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func main() {
	tree := &t.Tree{Value: 0,
		Left: &t.Tree{Value: -8,
			Left:  &t.Tree{Value: -10},
			Right: &t.Tree{Value: -6},
		},
		Right: &t.Tree{Value: 100},
	}
	fmt.Println(isTreeBF(tree))
}

func isTreeBF(tree *t.Tree) bool {
	return helper(tree, tree.LEFT()-1, tree.RIGHT()+1)
}

func helper(tree *t.Tree, left int, right int) bool {
	if tree == nil {
		return true
	}
	return tree.Value > left && tree.Value < right &&
		helper(tree.Left, left, tree.Value) &&
		helper(tree.Right, tree.Value, right)
}
