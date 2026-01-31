package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1, arr2 := []int{1, 4, 7, -1, 9, 0}, []int{5, 6, 1, 7, 8, 4}
	fmt.Println(esSort(arr1, arr2))
}

func esSort(arr1 []int, arr2 []int) []int {
	var res, pre []int
	var ex [][]int
	c := 0
	sort.Ints(arr1)
	sort.Ints(arr2)
	l, r := 0, 0
	for l < len(arr1) && r < len(arr2) {
		if arr1[l] == arr2[r] {
			c++
			pre = append(pre, arr1[l])
			ex = append(ex, pre)
			res = append(res, c)
			l, r = l+1, r+1
		} else if arr1[l] < arr2[r] {
			l++
			ex = append(ex, pre)
			res = append(res, c)
		} else {
			r++
		}
	}
	for l < r {
		l++
		ex = append(ex, pre)
		res = append(res, c)
	}

	fmt.Println(ex)
	return res
}
