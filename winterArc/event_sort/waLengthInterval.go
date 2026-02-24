package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{0, 3}, {5, 10}, {2, 4}, {6, 8}, {11, 12}, {16, 21}, {12, 19}, {20, 23}}
	fmt.Println(lenInter(arr))
}

func lenInter(arr [][]int) int {
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

	cnt := 0
	sum := 0
	l := 0
	for i := range st {
		sum += st[i][1]
		if sum == 0 {
			cnt += st[i][0] - st[l][0]
			l = i + 1
		}
	}
	return cnt
}
