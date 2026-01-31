package main

import (
	"fmt"
	"sort"
)

func Hotel(arr [][]int) int {
	n := len(arr)
	st := make([][]int, 2*n)
	index := 0
	for i := 0; i < n; i++ {
		st[index] = []int{arr[i][0], 1}
		index++
		st[index] = []int{arr[i][1], -1}
		index++
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})
	maxnum := 0
	cnt := 0
	for i := range st {
		cnt += st[i][1]
		maxnum = max(maxnum, cnt)
	}
	return maxnum
}

func main() {
	h := [][]int{{0, 5}, {2, 3}, {4, 6}, {10, 18}, {7, 12}, {11, 18}, {18, 21}}
	fmt.Println(Hotel(h))
}
