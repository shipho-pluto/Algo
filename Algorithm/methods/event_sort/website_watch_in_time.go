package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{0, 3}, {5, 10}, {2, 4}, {6, 8}, {11, 20}, {16, 21}, {12, 19}, {20, 23}}
	boss := []int{1, 2, 10, 18}
	fmt.Println(watchViewersInTime(arr, boss))
}

func watchViewersInTime(arr [][]int, boss []int) []int {
	n, b := len(arr), len(boss)
	idx := 0
	st := make([][]int, 2*n+b)
	for i := range n {
		st[idx] = []int{arr[i][0] - 1, 1}
		idx++
		st[idx] = []int{arr[i][1] + 1, -1}
		idx++
	}
	for i := range b {
		st[idx] = []int{boss[i], 0}
		idx++
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	res := make([]int, 0, b)
	cnt := 0
	for i := range st {
		cnt += st[i][1]
		if st[i][1] == 0 {
			res = append(res, cnt)
		}
	}
	return res
}
