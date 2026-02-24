package main

import (
	t "Algoritm/winterArc/trees/struct"
	"fmt"
)

func main() {
	tree := &t.TreeInt{Value: 5,
		Right: &t.TreeInt{Value: 8,
			Right: &t.TreeInt{Value: 9},
			Left:  &t.TreeInt{Value: 6},
		},
		Left: &t.TreeInt{Value: 2,
			Right: &t.TreeInt{Value: 4},
			Left:  &t.TreeInt{Value: 1},
		},
	}

	fmt.Println(bsTree(tree))
}

func minTree(t *t.TreeInt) int {
	if t.Left == nil {
		return t.Value
	}
	return minTree(t.Left)
}

func maxTree(t *t.TreeInt) int {
	if t.Right == nil {
		return t.Value
	}
	return maxTree(t.Right)
}

func bsTree(tree *t.TreeInt) bool {
	mnT := minTree(tree) - 1
	mxT := maxTree(tree) + 1
	return walkForTree2(mnT, mxT, tree)
}

func walkForTree2(mnT int, mxT int, t *t.TreeInt) bool {
	if t != nil {
		return t.Value < mxT && t.Value > mnT &&
			walkForTree2(mnT, t.Value, t.Left) &&
			walkForTree2(t.Value, mxT, t.Right)
	}
	return true
}
