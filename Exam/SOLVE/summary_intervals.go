package main

import (
	"fmt"
	"sort"
)

func summaryIntervals(i1, i2 [][]int) [][]int {
	n := len(i1)
	var res [][]int
	st := make([][]int, 4*n)
	in := 0
	for i := 0; i < n; i++ {
		st[in] = []int{i1[i][0], 1}
		in++
		st[in] = []int{i2[i][0], 1}
		in++
		st[in] = []int{i1[i][1], -1}
		in++
		st[in] = []int{i2[i][1], -1}
		in++
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	in = 0
	cnt := 0
	for i := range st {
		cnt += st[i][1]
		if cnt == 0 {
			res = append(res, []int{st[in][0], st[i][0]})
			in = i + 1
		}
	}
	return res
}

func main() {
	a1, a2 := [][]int{{1, 5}, {14, 20}, {23, 24}}, [][]int{{4, 6}, {10, 14}, {17, 19}}
	fmt.Println(summaryIntervals(a1, a2))
}
