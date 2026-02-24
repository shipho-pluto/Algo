package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{0, 3}, {5, 10}, {2, 4}, {6, 8}, {11, 20}, {16, 21}, {12, 19}, {20, 23}}
	time := []int{1, 2, 10, 18}
	fmt.Println(cntInTime(arr, time))
}

func cntInTime(arr [][]int, t []int) []int {
	n := len(arr) + len(t)
	st := make([][]int, 0, 2*n)
	for i := range arr {
		st = append(st, []int{arr[i][0], 1})
		st = append(st, []int{arr[i][1], -1})
	}
	for i := range t {
		st = append(st, []int{t[i], 0})
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] > st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	res := make([]int, 0, len(t))
	sum := 0
	for i := range st {
		sum += st[i][1]
		if st[i][1] == 0 {
			res = append(res, sum)
		}
	}

	return res
}
