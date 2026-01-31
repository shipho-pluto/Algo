package main

import (
	"Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func main() {
	arr := []int{1, 0, 0, 1}
	k := 2
	fmt.Println(maxOneWithKZero(arr, k))
	arr = []int{0, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1}
	fmt.Println(maxOneWithKZero(arr, k))
}

func is(q bool) int {
	if q {
		return 1
	}
	return 0
}

func maxOneWithKZero(arr []int, k int) int {
	dq := Structs.DeQueue{}
	n := len(arr)
	if n == k {
		return n
	}
	k += is(arr[0] == 0)
	l, r := 0, is(arr[0] == 0)
	for ; r < n; r++ {
		if arr[r] == 0 {
			k--
			dq.PutBack(r)
			if k < 0 {
				break
			}
		}
	}
	res := r - is(arr[0] == 0)
	for r = r + 1; r < n; r++ {
		if arr[r] == 0 {
			dq.PutBack(r)
			l = dq.TopFront()
			dq.PopFront()
			res = max(res, r-l-1)
		}
	}
	l = dq.TopFront()
	dq.PopFront()
	res = max(res, r-l-1)
	return res
}
