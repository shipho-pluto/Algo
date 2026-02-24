package main

import (
	t "Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func isSymmetric(left, right *t.Tree) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	if left.Value != right.Value {
		return false
	}
	return isSymmetric(left.Left, right.Right) && isSymmetric(left.Right, right.Left)
}

func main() {
	tree := &t.Tree{Value: 1,
		Right: &t.Tree{Value: 2,
			Right: &t.Tree{Value: 3,
				Right: &t.Tree{Value: 10},
				Left:  &t.Tree{Value: 9},
			},
			Left: &t.Tree{Value: 4,
				Right: &t.Tree{Value: 9},
				Left:  &t.Tree{Value: 10},
			},
		},
		Left: &t.Tree{Value: 2,
			Right: &t.Tree{Value: 4,
				Right: &t.Tree{Value: 10},
				Left:  &t.Tree{Value: 9},
			},
			Left: &t.Tree{Value: 3,
				Right: &t.Tree{Value: 9},
				Left:  &t.Tree{Value: 10},
			},
		},
	}

	fmt.Println(isSymmetric(tree.Left, tree.Right))
}
