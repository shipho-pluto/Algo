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
					Value: 80,
					Left: &t.TreeInt{
						Value: 50,
					},
					Right: &t.TreeInt{
						Value: 60,
						Left: &t.TreeInt{
							Value: 43,
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

	tree = &t.TreeInt{Value: 5,
		Right: &t.TreeInt{Value: 8},
		Left: &t.TreeInt{Value: 2,
			Right: &t.TreeInt{Value: 4},
			Left:  &t.TreeInt{Value: 1},
		},
	}

	fmt.Println(alvTree(tree))

	root := &t.TreeCh{Value: 'a',
		Left: &t.TreeCh{Value: 'b',
			Left: &t.TreeCh{Value: 'c'},
			Right: &t.TreeCh{Value: 'd',
				Right: &t.TreeCh{Value: 'd'},
			},
		},
		Right: &t.TreeCh{Value: 'a',
			Left: &t.TreeCh{Value: 'b',
				Left:  &t.TreeCh{Value: 'c'},
				Right: &t.TreeCh{Value: 'd'},
			},
		},
	}

	fmt.Println(twoEqTrees(root))

	tree = &t.TreeInt{Value: -1,
		Left: &t.TreeInt{Value: 1,
			Left: &t.TreeInt{Value: -10},
			Right: &t.TreeInt{Value: -6,
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
	fmt.Println(maxSumTree(tree))

	tree = &t.TreeInt{Value: 10,
		Left: &t.TreeInt{Value: 5,
			Left:  &t.TreeInt{Value: -2},
			Right: &t.TreeInt{Value: 7},
		},
		Right: &t.TreeInt{Value: 11,
			Right: &t.TreeInt{Value: 15},
		},
	}

	k := 3
	fmt.Println(findKMin(tree, &k))
	k = 3
	fmt.Println(findKMax(tree, &k))

	tree = &t.TreeInt{Value: 5,
		Right: &t.TreeInt{Value: 8,
			Right: &t.TreeInt{Value: 9},
			Left:  &t.TreeInt{Value: 6},
		},
		Left: &t.TreeInt{Value: 2,
			Right: &t.TreeInt{Value: 3},
			Left:  &t.TreeInt{Value: 1},
		},
	}

	fmt.Println(bsTree(tree))

	tree = &t.TreeInt{Value: 1,
		Right: &t.TreeInt{Value: 2,
			Right: &t.TreeInt{Value: 3,
				Right: &t.TreeInt{Value: 10},
				Left:  &t.TreeInt{Value: 9},
			},
			Left: &t.TreeInt{Value: 4,
				Right: &t.TreeInt{Value: 8},
				Left:  &t.TreeInt{Value: 7},
			},
		},
		Left: &t.TreeInt{Value: 2,
			Right: &t.TreeInt{Value: 4,
				Right: &t.TreeInt{Value: 7},
				Left:  &t.TreeInt{Value: 8},
			},
			Left: &t.TreeInt{Value: 3,
				Right: &t.TreeInt{Value: 9},
				Left:  &t.TreeInt{Value: 10},
			},
		},
	}

	fmt.Println(isTreeSymmetric(tree))
}

func isTreeSymmetric(tree *t.TreeInt) any {
	return walk6(tree.Left, tree.Right)
}

func walk6(l *t.TreeInt, r *t.TreeInt) bool {
	if l != nil && r != nil {
		if l.Value != r.Value {
			return false
		}
		return walk6(l.Left, r.Right) && walk6(l.Right, r.Left)
	} else if l == nil && r == nil {
		return true
	}
	return false
}

func maxT(tree *t.TreeInt) int {
	if tree.Right != nil {
		return maxT(tree.Right)
	}
	return tree.Value
}

func minT(tree *t.TreeInt) int {
	if tree.Left != nil {
		return minT(tree.Left)
	}
	return tree.Value
}

func bsTree(tree *t.TreeInt) any {
	mx := maxT(tree) + 1
	mn := minT(tree) - 1
	return walk5(tree, mx, mn)
}

func walk5(tree *t.TreeInt, mx int, mn int) bool {
	if tree != nil {
		return mn < tree.Value && tree.Value < mx &&
			walk5(tree.Left, tree.Value, mn) &&
			walk5(tree.Right, mx, tree.Value)
	}
	return true
}

func findKMax(tree *t.TreeInt, i *int) any {
	if tree == nil {
		return 0
	}

	res := findKMax(tree.Right, i)
	if res != 0 {
		return res
	}

	*i--
	if *i == 0 {
		return tree.Value
	}
	return findKMax(tree.Left, i)
}

func findKMin(tree *t.TreeInt, i *int) any {
	if tree == nil {
		return 0
	}
	res := findKMin(tree.Left, i)
	if res != 0 {
		return res
	}

	*i--
	if *i == 0 {
		return tree.Value
	}

	return findKMin(tree.Right, i)
}

func maxSumTree(tree *t.TreeInt) (any, int) {
	st := make(map[int]*t.TreeInt)
	_ = walk4(&st, tree)
	maxS := -1000000
	res := &t.TreeInt{}
	for v := range st {
		if v > maxS {
			maxS = v
			res = st[v]
		}
	}

	return res.Value, maxS
}

func walk4(m *map[int]*t.TreeInt, tree *t.TreeInt) int {
	if tree == nil {
		return 0
	}

	ans := tree.Value
	ls, rs := walk4(m, tree.Left), walk4(m, tree.Right)
	if ls > 0 {
		ans += ls
	}
	if rs > 0 {
		ans += rs
	}

	if (*m)[ans] == nil {
		(*m)[ans] = tree
	}

	return ans
}

func twoEqTrees(root *t.TreeCh) (any, int) {
	st := make(map[int][2]*t.TreeCh)
	_ = walk3(&st, root)
	res := [2]string{}
	maxH := 0
	for _, v := range st {
		h1, h2 := maxHighC(v[0]), maxHighC(v[1])
		if h1+h2 > maxH {
			maxH = h1 + h2
			res = [2]string{string(v[0].Value), string(v[1].Value)}
		}
	}

	return res, maxH
}

func maxHighC(ch *t.TreeCh) int {
	if ch != nil {
		return max(maxHighC(ch.Left), maxHighC(ch.Right)) + 1
	}
	return 0
}

func walk3(m *map[int][2]*t.TreeCh, root *t.TreeCh) int {
	if root == nil {
		return 0
	}

	ans := 1 << (root.Value - 'a')
	ans |= walk3(m, root.Left)
	ans |= walk3(m, root.Right)

	node := (*m)[ans]
	if node[0] == nil {
		node[0] = root
	} else if node[1] == nil {
		node[1] = root
	}
	(*m)[ans] = node

	return ans
}

func ABS(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func alvTree(tree *t.TreeInt) bool {
	return ABS(maxHigh(tree.Left)-maxHigh(tree.Right)) <= 1
}

func maxHighVal(tree *t.TreeInt) any {
	maxEl := 0
	maxH := maxHigh(tree)
	return walk1(tree, &maxEl, 1, maxH)
}

func walk1(tree *t.TreeInt, v *int, h int, H int) any {
	if tree != nil {
		if h == H {
			*v = max(*v, tree.Value)
		}
		walk1(tree.Left, v, h+1, H)
		walk1(tree.Right, v, h+1, H)
	}
	return *v
}

func maxHigh(root *t.TreeInt) int {
	if root != nil {
		return max(maxHigh(root.Left), maxHigh(root.Right)) + 1
	}
	return 0
}
