package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{0, 3}, {5, 10}, {2, 4}, {6, 8}, {11, 12}, {16, 21}, {12, 19}, {20, 23}}
	fmt.Println(maxParallel(arr))
}

func maxParallel(arr [][]int) int {
	n := len(arr)
	st := make([][]int, 0, 2*n)
	for i := range arr {
		st = append(st, []int{arr[i][0], 1})
		st = append(st, []int{arr[i][1], -1})
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	maxP := 0
	c := 0
	for i := range st {
		c += st[i][1]
		maxP = max(maxP, c)
	}

	return maxP
}
