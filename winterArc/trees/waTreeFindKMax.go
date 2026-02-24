package main

import (
	t "Algoritm/winterArc/trees/struct"
	"fmt"
)

func main() {
	tree := &t.TreeInt{Value: 10,
		Left: &t.TreeInt{Value: 5,
			Left:  &t.TreeInt{Value: -2},
			Right: &t.TreeInt{Value: 7},
		},
		Right: &t.TreeInt{Value: 11,
			Right: &t.TreeInt{Value: 15},
		},
	}

	k := 3
	fmt.Println(findKMax(tree, &k))
}

func findKMax(tree *t.TreeInt, k *int) int {
	if tree == nil {
		return 0
	}

	res := findKMax(tree.Right, k)
	if res != 0 {
		return res
	}

	*k -= 1
	if *k == 0 {
		return tree.Value
	}
	return findKMax(tree.Left, k)
}
