package main

import (
	"fmt"
	"sort"
)

func main() {
	in := [][]int{{0, 2}, {4, 9}, {3, 5}, {10, 11}, {14, 19}, {10, 14}}
	fmt.Println(sumInter(in))
}

func sumInter(arr [][]int) [][]int {
	n := len(arr)
	st := make([][]int, 0, 2*n)
	for i := range arr {
		st = append(st, []int{arr[i][0], 1})
		st = append(st, []int{arr[i][1], -1})
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] > st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	var res [][]int
	l := 0
	sum := 0
	for i := range st {
		sum += st[i][1]
		if sum == 0 {
			res = append(res, []int{st[l][0], st[i][0]})
			l = i + 1
		}
	}

	return res
}
