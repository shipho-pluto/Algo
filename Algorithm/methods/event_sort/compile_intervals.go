package main

import (
	"fmt"
	"sort"
)

func main() {
	in := [][]int{{0, 2}, {4, 9}, {3, 5}, {10, 11}, {14, 19}, {10, 14}}
	fmt.Println(compileIntervals(in))
}

func compileIntervals(in [][]int) [][]int {
	n := len(in)
	st := make([][]int, 2*n)
	idx := 0
	for i := range n {
		st[idx] = []int{in[i][0], 1}
		idx++
		st[idx] = []int{in[i][1], -1}
		idx++
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] > st[j][1]
		}
		return st[i][0] < st[j][0]
	})
	fmt.Println(st)

	var res [][]int
	l := 0
	cnt := 0
	for r := range st {
		cnt += st[r][1]
		if cnt == 0 {
			res = append(res, []int{st[l][0], st[r][0]})
			l = r + 1
		}
	}

	return res
}
