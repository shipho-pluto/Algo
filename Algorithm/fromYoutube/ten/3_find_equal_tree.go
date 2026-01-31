package main

type tree struct {
	value rune
	left  *tree
	right *tree
}

func walkForTree(st *map[int][2]*tree, t *tree) int {
	if t == nil {
		return 0
	}
	ans := 1 << (t.value - 'a')
	ans |= walkForTree(st, t.left)
	ans |= walkForTree(st, t.right)

	node := (*st)[ans]
	if node[0] == nil {
		node[0] = t
	} else if node[1] == nil {
		node[1] = t
	}
	(*st)[ans] = node

	return ans
}

func findEqualTree(root *tree) [2]*tree {
	var res [2]*tree
	st := make(map[int][2]*tree)
	_ = walkForTree(&st, root)
	maxHigh := 0
	for _, node := range st {
		h1, h2 := high(node[0]), high(node[1])
		if h1+h2 > maxHigh {
			maxHigh = h1 + h2
			res = node
		}
	}
	return res
}

func high(t *tree) int {
	if t == nil {
		return 0
	}
	return max(high(t.left), high(t.right)) + 1
}
