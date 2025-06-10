package main

import "fmt"

type TNode5 struct {
	value rune
	left  *TNode5
	right *TNode5
}

func GO(st *map[int][2]*TNode5, t *TNode5) int {
	if t == nil {
		return 0
	}
	ans := 1 << (t.value - 'A')

	ans |= GO(st, t.left)
	ans |= GO(st, t.right)

	nodes := (*st)[ans]
	if nodes[0] == nil {
		nodes[0] = t
	} else if nodes[1] == nil {
		nodes[1] = t
	}
	(*st)[ans] = nodes

	return ans
}

func FindTwoNode(t *TNode5) [2]*TNode5 {
	st := make(map[int][2]*TNode5)
	_ = GO(&st, t)

	ans := [2]*TNode5{nil, nil}
	for _, value := range st {
		if value[0] != nil && value[1] != nil {
			ans[0] = value[0]
			ans[1] = value[1]
		}
	}
	return ans
}

func main() {
	t1 := &TNode5{'B', &TNode5{'A', nil, nil}, &TNode5{'A', nil, nil}}
	fmt.Println(FindTwoNode(t1))
}
