package main

import "fmt"

type tree struct {
	val   rune
	left  *tree
	right *tree
}

func main() {
	t := &tree{
		val: 'A',
		left: &tree{
			val: 'D',
			left: &tree{
				val: 'A',
				left: &tree{
					val: 'B',
					left: &tree{
						val: 'C',
					},
				},
			},
		},
	}

	fmt.Println(symmetricTree(t))
}

func symmetricTree(t *tree) any {
	st := map[int][2]*tree{}
	_ = walkForTree(t, &st)
	res := [2]*tree{}
	maxHigh := 0

	for k, node := range st {
		high := cntOne(k) * 2
		if high > maxHigh && node[0] != nil && node[1] != nil {
			fmt.Println(k)
			maxHigh = high
			res = node
		}
	}
	fmt.Println(maxHigh)

	return res
}

func cntOne(k int) int {
	cnt := 0
	for k != 0 {
		cnt += k % 2
		k /= 2
	}
	return cnt
}

func walkForTree(t *tree, st *map[int][2]*tree) int {
	if t == nil {
		return 0
	}

	ans := 1 << (t.val - 'A')
	ans |= walkForTree(t.left, st)
	ans |= walkForTree(t.right, st)

	node := (*st)[ans]
	if node[0] == nil {
		node[0] = t
	} else if node[1] == nil {
		node[1] = t
	}
	(*st)[ans] = node

	return ans
}
