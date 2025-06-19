package main

import (
	"fmt"
	"sort"
)

func CompileIntervals(arr [][]int) [][]int {
	var res [][]int
	n := len(arr)
	st := make([][]int, 2*n)
	index := 0
	for _, v := range arr {
		st[index] = []int{v[0], 1}
		st[index+1] = []int{v[1], -1}
		index += 2
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] > st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	cnt := 0
	index = 0
	for i := 0; i < 2*n; i++ {
		cnt += st[i][1]
		if cnt == 0 {
			res = append(res, []int{st[index][0], st[i][0]})
			index = i + 1
		}

	}
	return res
}

func main() {
	arr := [][]int{{0, 3}, {5, 10}, {2, 4}, {6, 8}, {11, 12}, {16, 21}, {12, 19}, {20, 23}}
	fmt.Println(CompileIntervals(arr))
}
