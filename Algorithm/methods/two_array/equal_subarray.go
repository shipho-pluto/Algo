package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1, arr2 := []int{1, 3, -4, 0, 1}, []int{-1, 4, 6, 0, 1, 1}
	fmt.Println(esMap(arr1, arr2))
	arr1, arr2 = []int{1, 3, -4, 0, 1}, []int{0, -4}
	fmt.Println(es(arr1, arr2))
}

func es(arr1 []int, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var res []int
	n, m := len(arr1), len(arr2)
	for l, r := 0, 0; l < n && r < m; {
		if arr1[l] < arr2[r] {
			l++
		} else if arr1[l] == arr2[r] {
			res = append(res, arr1[l])
			l++
			r++
		} else {
			r++
		}
	}
	return res
}

func esMap(arr1 []int, arr2 []int) []int {
	var res []int
	st := make(map[int]int)
	for i := range arr1 {
		st[arr1[i]]++
	}
	for i := range arr2 {
		if _, ok := st[arr2[i]]; ok {
			res = append(res, arr2[i])
			st[arr2[i]]--
			if st[arr2[i]] == 0 {
				delete(st, arr2[i])
			}
		}
	}
	return res
}
