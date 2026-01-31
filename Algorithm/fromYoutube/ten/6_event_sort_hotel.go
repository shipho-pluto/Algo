package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{0, 1}, {2, 9}, {4, 6}, {5, 12}, {14, 18}, {10, 15}, {16, 22}, {20, 24}}
	fmt.Println(hotelMaxPeople(arr))
}

func hotelMaxPeople(arr [][]int) int {
	n := len(arr)
	st := make([][]int, 2*n)
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

	var res int
	cnt := 0
	for time := range st {
		cnt += st[time][1]
		res = max(res, cnt)
	}
	return res
}
