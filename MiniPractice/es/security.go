package main

import (
	"fmt"
	"sort"
)

func Security(arr [][]int) bool {
	n := len(arr)
	st := make([][]int, 2*n)
	for index, i := 0, 0; i < n; func() { i++; index += 2 }() {
		st[index] = []int{arr[i][0] + 1, 1}
		st[index+1] = []int{arr[i][1] + 1, -1}
	}
	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] > st[j][1]
		}
		return st[i][0] < st[j][0]
	})
	fmt.Println(st)
	cnt := 0
	index := 0
	for i := range st {
		cnt += st[i][1]
		index = st[i][0]
		if cnt == 0 && index < 10000 {
			return false
		}
	}
	return index >= 10000
}

func main() {
	n := [][]int{{0, 1}, {1, 10}, {2, 8}, {11, 12}, {12, 10000}}
	fmt.Println(Security(n))
}
