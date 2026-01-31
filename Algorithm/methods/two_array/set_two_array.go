package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1, arr2 := []int{1, 4, 5, 8, 10}, []int{3, 5, 8} // 1, 3, 4, 10
	fmt.Println(staSort(arr1, arr2))

	arr1, arr2 = []int{1, 1, 1, 1}, []int{1, 1, 1} // 8, 4, 2, 7, 5
	fmt.Println(staMap(arr1, arr2))
}

func staSort(arr1 []int, arr2 []int) []int {
	var res []int
	sort.Ints(arr1)
	sort.Ints(arr2)
	for l, r := 0, 0; l < len(arr1) || r < len(arr2); {
		if l < len(arr1) && r < len(arr2) {
			if arr1[l] < arr2[r] {
				res = append(res, arr1[l])
				l++
			} else if arr1[l] == arr2[r] {
				l, r = l+1, r+1
			} else {
				res = append(res, arr2[r])
				r++
			}
		} else if l < len(arr1) {
			res = append(res, arr1[l])
			l++
		} else {
			res = append(res, arr2[r])
			r++
		}
	}
	return res
}

func staMap(arr1 []int, arr2 []int) []int {
	var res []int
	st := make(map[int]int)
	for i := range arr1 {
		st[arr1[i]]++
	}
	for i := range arr2 {
		if _, ok := st[arr2[i]]; ok {
			st[arr2[i]]--
			if st[arr2[i]] == 0 {
				delete(st, arr2[i])
			}
		} else {
			res = append(res, arr2[i])
		}
	}

	for i := range st {
		res = append(res, i)
	}

	return res
}
