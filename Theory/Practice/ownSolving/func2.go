package main

import "fmt"

type BTree2 struct {
	value rune
	left  *BTree2
	right *BTree2
}

func Height(t *BTree2, loop int) int {
	if t != nil {
		loop = max(Height(t.right, loop+1), Height(t.left, loop+1))
	}
	return loop
}

func GO(st *map[int][2]*BTree2, t *BTree2) int {
	if t == nil {
		return 0
	}
	ans := 1 << (t.value - 'a')

	ans |= GO(st, t.left)
	ans |= GO(st, t.right)

	note := (*st)[ans]
	if note[0] == nil {
		note[0] = t
	} else if note[1] == nil {
		note[1] = t
	}
	(*st)[ans] = note

	return ans
}

func Find2EqualTree(t *BTree2) [2]*BTree2 {
	st := make(map[int][2]*BTree2)
	maxLoop := 0
	_ = GO(&st, t)
	ans := [2]*BTree2{nil, nil}
	for _, val := range st {
		if val[0] != nil && val[1] != nil {
			h1, h2 := Height(val[0], 0), Height(val[1], 0)
			if h1+h2 > maxLoop {
				maxLoop = h1 + h2
				ans[0] = val[0]
				ans[1] = val[1]
			}
		}
	}

	return ans
}

func main() {
	T := &BTree2{'a', &BTree2{'b', &BTree2{'c', nil, nil}, nil}, &BTree2{'c', nil, nil}}
	fmt.Println(Find2EqualTree(T))
}
