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

	node := (*st)[ans]
	if node[0] == nil {
		node[0] = t
	} else if node[1] == nil {
		node[1] = t
	}
	(*st)[ans] = node

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
			break
		}
	}
	return ans
}

func main() {
	t1 := &TNode5{'B', &TNode5{'A', nil, nil}, &TNode5{'A', nil, nil}}
	fmt.Println(FindTwoNode(t1))
}
