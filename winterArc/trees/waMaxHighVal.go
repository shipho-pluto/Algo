package main

import (
	t "Algoritm/winterArc/trees/struct"
	"fmt"
)

func main() {
	tree := &t.TreeInt{
		Value: 10,
		Right: &t.TreeInt{
			Value: 20,
			Right: &t.TreeInt{
				Value: 30,
				Right: &t.TreeInt{
					Value: 40,
					Left: &t.TreeInt{
						Value: 50,
					},
					Right: &t.TreeInt{
						Value: 60,
						Left: &t.TreeInt{
							Value: 80,
						},
						Right: &t.TreeInt{
							Value: 31,
						},
					},
				},
			},
			Left: &t.TreeInt{
				Value: 15,
			},
		},
	}

	fmt.Println(maxHighVal(tree))
}

func maxHighVal(tree *t.TreeInt) (int, int) {
	maxEl := tree.Value
	maxH := 0
	highVal(tree, &maxEl, &maxH, 1)
	return maxH, maxEl
}

func highVal(tree *t.TreeInt, m *int, h *int, loop int) {
	if tree != nil {
		if tree.Value >= *m && loop >= *h {
			*m = tree.Value
			*h = loop
		}
		highVal(tree.Right, m, h, loop+1)
		highVal(tree.Left, m, h, loop+1)
	}
}
