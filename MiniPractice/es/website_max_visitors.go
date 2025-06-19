package main

import (
	"fmt"
	"sort"
)

func WebsiteMaxVisitors(arr [][]int) int {
	n := len(arr)
	m := make([][]int, 2*n)
	for i, in := 0, 0; i < n; func() { i++; in += 2 }() {
		m[in] = []int{arr[i][0], 1}
		m[in+1] = []int{arr[i][1], -1}
	}

	sort.Slice(m, func(i, j int) bool {
		if m[i][0] == m[j][0] {
			return m[i][1] < m[j][1]
		}
		return m[i][0] < m[j][0]
	})

	res, cnt := 0, 0
	for _, v := range m {
		cnt += v[1]
		res = max(res, cnt)
	}
	return res

}

func main() {
	arr := [][]int{{0, 3}, {5, 10}, {2, 4}, {6, 8}, {11, 12}, {16, 21}, {12, 19}, {20, 23}}
	fmt.Println(WebsiteMaxVisitors(arr))
}
