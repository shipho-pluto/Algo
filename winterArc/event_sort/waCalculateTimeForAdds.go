package main

import (
	"fmt"
	"sort"
)

func main() {
	n := [][]int{{1, 4}, {4, 6}, {5, 6}, {8, 10}, {6, 9}, {7, 9}, {11, 17}, {10, 14}, {11, 15}, {12, 13}}
	fmt.Println(calculateAdds(n, 1))
}

func calculateAdds(arr [][]int, k int) (int, int, int) {
	var st [][]int
	for i := range arr {
		if arr[i][1]-arr[i][0] >= k {
			st = append(st, []int{arr[i][0], 1})
			st = append(st, []int{arr[i][1], -1})
		}
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	res := [][]int{{0, 0}, {0, 0}}
	max1 := 0
	max2 := 0
	view := 0
	for i := range st {
		view += st[i][1]
		if view > max1 {
			max2 = max1
			max1 = view
			res = [][]int{{st[i][0], max1}, {res[0][0], max2}}
		} else if view > max2 {
			max2 = view
			res[1] = []int{st[i][0], max2}
		}
	}

	return res[0][0], res[1][0], res[0][1] + res[1][1]
}
