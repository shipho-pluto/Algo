package main

import (
	t "Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func main() {
	tree := &t.Tree{Value: 10,
		Left: &t.Tree{Value: 5,
			Left: &t.Tree{Value: -2,
				Left:  nil,
				Right: nil,
			},
			Right: &t.Tree{Value: 7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &t.Tree{Value: 11,
			Right: &t.Tree{Value: 15,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(findKMaxElement(tree, 3))
}

func findKMaxElement(tree *t.Tree, k int) any {
	return wrapperMax(tree, &k)
}

func wrapperMax(tree *t.Tree, k *int) int {
	if tree == nil {
		return 0
	}

	res := wrapperMax(tree.Right, k)
	if res != 0 {
		return res
	}

	(*k)--
	if (*k) == 0 {
		return tree.Value
	}

	return wrapperMax(tree.Left, k)
}
