package main

import (
	t "Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func main() {
	tree := &t.Tree{Value: 10,
		Left: &t.Tree{Value: 5,
			Left:  &t.Tree{Value: -2},
			Right: &t.Tree{Value: 7},
		},
		Right: &t.Tree{Value: 11,
			Right: &t.Tree{Value: 15},
		},
	}
	fmt.Println(findKMinElement(tree, 3))
}

func findKMinElement(tree *t.Tree, k int) int {
	return wrapperMin(tree, &k)
}

func wrapperMin(tree *t.Tree, k *int) int {
	if tree == nil {
		return 0
	}

	res := wrapperMin(tree.Left, k)
	if res != 0 {
		return res
	}

	*k -= 1
	if *k == 0 {
		return tree.Value
	}

	return wrapperMin(tree.Right, k)
}
