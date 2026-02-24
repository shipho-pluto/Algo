package main

import (
	t "Algoritm/winterArc/trees/struct"
	"fmt"
)

func main() {
	tree := &t.TreeInt{Value: 0,
		Left: &t.TreeInt{Value: -8,
			Left: &t.TreeInt{Value: -10,
				Right: &t.TreeInt{Value: -11},
			},
			Right: &t.TreeInt{Value: -7,
				Right: &t.TreeInt{Value: -6},
			},
		},
		Right: &t.TreeInt{Value: 100,
			Left: &t.TreeInt{Value: 8,
				Left:  &t.TreeInt{Value: 10},
				Right: &t.TreeInt{Value: 6},
			},
		},
	}

	t1, t2 := 8, 10

	fmt.Println(nearParent(tree, t1, t2))
	fmt.Println(nearParent2(tree, t1, t2))
}

func nearParent2(tree *t.TreeInt, t1 int, t2 int) any {
	var ans []string
	walkForTree32(&ans, tree, t1, t2, "")
	i := 0
	for ; i < min(len(ans[0]), len(ans[1])) && ans[0][i] == ans[1][i]; i++ {
	}
	res := ans[0][:i]
	for j := range res {
		if res[j] == '1' {
			tree = tree.Right
		} else {
			tree = tree.Left
		}
	}
	return tree.Value
}

func walkForTree32(ans *[]string, tree *t.TreeInt, t1 int, t2 int, loop string) {
	if tree != nil {
		if tree.Value == t1 || tree.Value == t2 {
			*ans = append(*ans, loop)
			if len(*ans) == 2 {
				return
			}
		}
		walkForTree32(ans, tree.Left, t1, t2, loop+"0")
		walkForTree32(ans, tree.Right, t1, t2, loop+"1")
	}
}

func nearParent(tree *t.TreeInt, t1 int, t2 int) any {
	var st [][]int
	walkForTree3(&st, tree, 0)
	cnt := 0
	res := st[0]
	for l, r := 0, 1; l < len(st); r++ {
		if r < len(st) && (st[r][0] == t1 || st[r][0] == t2) && st[r][1] > st[l][1] {
			cnt++
		}
		if r < len(st) && st[r][1] <= st[l][1] || r == len(st) {
			if cnt == 2 && st[l][1] >= res[1] {
				res = st[l]
				l, r = l+1, l+1
			} else if cnt == 1 {
				return res[0]
			} else {
				l = r
			}
			cnt = 0
			if l < len(st) && (st[l][0] == t1 || st[l][0] == t2) {
				cnt++
			}
		}
	}
	return res[0]
}

func walkForTree3(st *[][]int, tree *t.TreeInt, loop int) {
	if tree != nil {
		*st = append(*st, []int{tree.Value, loop})
		walkForTree3(st, tree.Left, loop+1)
		walkForTree3(st, tree.Right, loop+1)
	}
}
