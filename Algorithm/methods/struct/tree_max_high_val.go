package main

import (
	t "Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func findHighest(t *t.Tree, loop int) [2]int {
	if t == nil {
		return [2]int{}
	}
	if t.Right == nil && t.Left == nil {
		return [2]int{t.Value, loop + 1}
	}
	l, r := findHighest(t.Left, loop+1), findHighest(t.Right, loop+1)
	if l[1] == r[1] && l[0] > r[0] || l[1] > r[1] {
		return l
	}
	return r

}

func main() {
	tree := &t.Tree{
		Value: 10,
		Right: &t.Tree{
			Value: 20,
			Right: &t.Tree{
				Value: 30,
				Right: &t.Tree{
					Value: 40,
					Left: &t.Tree{
						Value: 50,
					},
					Right: &t.Tree{
						Value: 60,
						Left: &t.Tree{
							Value: 80,
						},
						Right: &t.Tree{
							Value: 31,
						},
					},
				},
			},
			Left: &t.Tree{
				Value: 15,
			},
		},
	}
	fmt.Println(findHighest(tree, 0))
}
