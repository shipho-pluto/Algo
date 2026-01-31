package main

import (
	"fmt"
	"sort"
)

func main() {
	n := [][]int{{1, 4}, {4, 6}, {5, 6}, {8, 10}, {6, 9}, {7, 9}, {11, 17}, {10, 14}, {11, 15}, {12, 13}}
	fmt.Println(calculateAdvertisement(n, 1))
}

func calculateAdvertisement(arr [][]int, k int) (int, int, int) {
	n := len(arr)
	var st [][]int
	for i := range n {
		if arr[i][1]-arr[i][0] >= k {
			st = append(st, []int{arr[i][0], 1})
			st = append(st, []int{arr[i][1], -1})
		}
	}
	m := len(st)

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	pref := make([][]int, m)
	cnt := 0
	for i := 0; i < m; i++ {
		cnt += st[i][1]
		pref[i] = []int{st[i][0], cnt}
	}

	var view [][]int
	for i := 1; i < m; i++ {
		if pref[i][1] < pref[i-1][1] {
			view = append(view, pref[i-1])
			for ; i < m; i++ {
				if pref[i][1] == 0 {
					break
				}
			}
		}
	}

	res := [][]int{{0, 0}, {0, 0}}
	for i := 0; i < len(view); i++ {
		if view[i][1] > res[0][1] {
			res[1] = res[0]
			res[0] = view[i]
		} else if view[i][1] > res[1][1] {
			res[1] = view[i]
		}
	}
	return res[1][0], res[0][0], res[0][1] + res[1][1]
}
