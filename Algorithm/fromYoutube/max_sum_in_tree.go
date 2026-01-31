package main

import "fmt"

type TreeStruct struct {
	Val         int
	Left, Right *TreeStruct
}

func main() {
	t := &TreeStruct{Val: -100,
		Left: &TreeStruct{Val: -200,
			Right: &TreeStruct{Val: -3,
				Left: &TreeStruct{Val: -4},
			},
		},
		Right: &TreeStruct{Val: -5,
			Left: &TreeStruct{Val: -1000,
				Left: &TreeStruct{Val: -7,
					Right: &TreeStruct{Val: -8},
				},
			},
		},
	}
	fmt.Println(maxSumTree(t))
}

func maxSumTree(t *TreeStruct) int {
	st := make(map[int]bool)
	_ = walk(&st, t)
	res := t.Val
	for k := range st {
		res = max(res, k)
	}
	return res
}

func walk(st *map[int]bool, t *TreeStruct) int {
	if t == nil {
		return 0
	}

	ans := t.Val
	ansL := walk(st, t.Left)
	ansR := walk(st, t.Right)
	if ans+ansL > ans {
		ans += ansL
	}
	if ans+ansR > ans {
		ans += ansR
	}

	(*st)[ans] = true

	return ans
}
