package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1, arr2 := []int{1, 3, -4, 0, 1}, []int{-1, 4, 6, 0, 1}
	fmt.Println(equalSubarray(arr1, arr2))
	arr1, arr2 = []int{1, 3, -4, 0, 0, 1}, []int{0, -4, 0}
	fmt.Println(equalSubarrayLR(arr1, arr2))
}

func equalSubarray(arr1 []int, arr2 []int) []int {
	st := make(map[int]int)

	for i := range arr1 {
		st[arr1[i]]++
	}
	var res []int
	for i := range arr2 {
		if _, ok := st[arr2[i]]; ok {
			st[arr2[i]]--
			res = append(res, arr2[i])
			if st[arr2[i]] == 0 {
				delete(st, arr2[i])
			}
		}
	}

	return res
}

func equalSubarrayLR(arr1 []int, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	var res []int
	for l, r := 0, 0; l < len(arr1) && r < len(arr2); {
		if arr1[l] < arr2[r] {
			l++
		} else if arr1[l] > arr2[r] {
			r++
		} else {
			res = append(res, arr1[l])
			l, r = l+1, r+1
		}
	}

	return res
}
