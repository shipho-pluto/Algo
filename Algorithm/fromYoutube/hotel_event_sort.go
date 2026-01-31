package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{1, 4}, {5, 9}, {3, 14}, {13, 14}, {15, 19}, {18, 21}, {20, 23}, {22, 23}, {21, 24}}
	fmt.Println(hotelSortEvent(arr))
}

func hotelSortEvent(arr [][]int) int {
	n := len(arr)
	st := make([][]int, 2*n)
	idx := 0
	for i := range n {
		st[idx] = []int{arr[i][0], 1}
		idx++
		st[idx] = []int{arr[i][1], -1}
		idx++
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	cnt := 0
	res := 0
	for i := range st {
		cnt += st[i][1]
		res = max(res, cnt)
	}
	return res
}
