package main

import "fmt"

type TreeNode struct {
	Val   rune
	Left  *TreeNode
	Right *TreeNode
}

func rec(st *map[int][2]*TreeNode, t *TreeNode) int {
	if t == nil {
		return 0
	}

	ans := 1 << (t.Val - 'A')

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
	t := &TreeNode{Val: 'A',
		Left: &TreeNode{Val: 'B',
			Right: &TreeNode{Val: 'C',
				Left: &TreeNode{Val: 'D'},
			},
		},
		Right: &TreeNode{Val: 'A',
			Left: &TreeNode{Val: 'B',
				Left: &TreeNode{Val: 'C',
					Right: &TreeNode{Val: 'D'},
				},
			},
		},
	}
	fmt.Println(findTwoEqTree(t))
}
