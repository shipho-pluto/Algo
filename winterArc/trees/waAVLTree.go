package main

import (
	t "Algoritm/winterArc/trees/struct"
	"fmt"
)

func main() {
	tree := &t.TreeInt{Value: 5,
		Right: &t.TreeInt{Value: 8},
		Left: &t.TreeInt{Value: 2,
			Right: &t.TreeInt{Value: 4},
			Left:  &t.TreeInt{Value: 1},
		},
	}

	fmt.Println(alvTree(tree))
}

func alvTree(t *t.TreeInt) bool {
	hL := highTree2(t.Left)
	hR := highTree2(t.Right)
	return hL == hR || hL-1 == hR || hR-1 == hL
}

func highTree2(t *t.TreeInt) int {
	if t != nil {
		return max(highTree2(t.Left)+1, highTree2(t.Right)+1)
	}
	return 1
}
