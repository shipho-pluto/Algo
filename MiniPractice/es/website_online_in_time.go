package main

import (
	"fmt"
	"sort"
)

func WebsiteOnlineVisitors(arr [][]int, boss []int) []int {
	n, l := len(arr), len(boss)
	m := make([][]int, 2*n+l)
	in := 0
	for i := 0; i < n; func() { i++; in += 2 }() {
		m[in] = []int{arr[i][0], 1}
		m[in+1] = []int{arr[i][1] + 1, -1}
	}
	for i := 0; i < l; func() { i++; in++ }() {
		m[in] = []int{boss[i], 2}
	}

	sort.Slice(m, func(i, j int) bool {
		if m[i][0] == m[j][0] {
			return m[i][1] < m[j][1]
		}
		return m[i][0] < m[j][0]
	})

	cnt, index := 0, 0
	res := make([]int, l)
	for _, v := range m {
		if v[1] == 2 {
			res[index] = cnt
			index++
		} else {
			cnt += v[1]
		}

	}
	return res
}

func main() {
	arr := [][]int{{0, 3}, {5, 10}, {2, 4}, {6, 8}, {11, 20}, {16, 21}, {12, 19}, {20, 23}}
	arr2 := []int{1, 2, 10, 18}
	fmt.Println(WebsiteOnlineVisitors(arr, arr2))
}
