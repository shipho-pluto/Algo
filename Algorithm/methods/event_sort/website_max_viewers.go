package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{0, 3}, {5, 10}, {2, 4}, {6, 8}, {11, 12}, {16, 21}, {12, 19}, {20, 23}}
	fmt.Println(websiteMaxViewers(arr))
}

func websiteMaxViewers(arr [][]int) int {
	n := len(arr)
	st := make([][]int, n*2)
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

	res := 0
	cnt := 0
	for i := range st {
		cnt += st[i][1]
		res = max(res, cnt)
	}
	return res
}
