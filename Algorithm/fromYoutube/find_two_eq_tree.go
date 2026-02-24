package main

import (
	"fmt"
)

type TreeNode struct {
	Value rune
	Left  *TreeNode
	Right *TreeNode
}

func rec(st *map[int][2]*TreeNode, t *TreeNode) int {
	if t == nil {
		return 0
	}

	ans := 1 << (t.Value - 'a')

	ans |= rec(st, t.Left)
	ans |= rec(st, t.Right)

	node := (*st)[ans]
	if node[0] == nil {
		node[0] = t
	} else if node[1] == nil {
		node[1] = t
	}
	(*st)[ans] = node

	return ans
}

func high(t *TreeNode) int {
	if t == nil {
		return 1
	}
	return max(high(t.Left), high(t.Right)) + 1
}

func findTwoEqTree(t *TreeNode) [2]*TreeNode {
	st := make(map[int][2]*TreeNode)
	res := [2]*TreeNode{}

	_ = rec(&st, t)
	maxHigh := 0
	for _, val := range st {
		if val[0] != nil && val[1] != nil {
			h1, h2 := high(val[0]), high(val[1])
			if h1+h2 > maxHigh {
				res = val
				maxHigh = h1 + h2
				fmt.Println(maxHigh)
			}
		}
	}

	return res
}

func main() {
	t := &TreeNode{Value: 'a',
		Left: &TreeNode{Value: 'b',
			Left: &TreeNode{Value: 't'},
			Right: &TreeNode{Value: 'u',
				Right: &TreeNode{Value: 'e'},
			},
		},
		Right: &TreeNode{Value: 'a',
			Left: &TreeNode{Value: 'b',
				Left:  &TreeNode{Value: 'c'},
				Right: &TreeNode{Value: 'd'},
			},
		},
	}
	fmt.Println(findTwoEqTree(t))
}
