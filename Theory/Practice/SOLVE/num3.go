package main

import "fmt"

type tree struct {
	val   rune
	left  *tree
	right *tree
}

func recursiveBuild(st *map[int][2]*tree, t *tree) int {
	if t == nil {
		return 0
	}

	ans := 1 << (t.val - 'A')
	ans |= recursiveBuild(st, t.left)
	ans |= recursiveBuild(st, t.right)

	node := (*st)[ans]
	if node[0] == nil {
		node[0] = t
	} else if node[1] == nil {
		node[1] = t
	}
	(*st)[ans] = node

	return ans
}

func height(t *tree) int {
	if t == nil {
		return 0
	}
	return 1 + max(height(t.left), height(t.right))
}

func FindTwoEqTree(t *tree) ([2]*tree, int) {
	st := make(map[int][2]*tree)
	_ = recursiveBuild(&st, t)

	res := [2]*tree{}
	maxDeep := 0
	for _, v := range st {
		if v[0] != nil && v[1] != nil {
			if maxDeep < height(v[0])+height(v[1]) {
				res[0] = v[0]
				res[1] = v[1]
				maxDeep = height(v[0]) + height(v[1])
			}
		}
	}

	return res, maxDeep
}

func main() {
	t := &tree{'A',
		&tree{'B',
			&tree{'C',
				nil,
				nil},
			nil},
		&tree{'C',
			&tree{'A',
				&tree{'B',
					&tree{'C',
						nil,
						nil},
					nil},
				&tree{'C',
					nil,
					nil}},
			nil}}
	fmt.Println(FindTwoEqTree(t))
}
