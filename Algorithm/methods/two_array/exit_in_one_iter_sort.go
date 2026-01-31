package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1, arr2 := []int{1, 4, 6, 10}, []int{4, 5, 6, 7}
	fmt.Println(eioSort(arr1, arr2))
}

func eioSort(arr1 []int, arr2 []int) []int {
	var res, pre []int
	var ex [][]int
	sort.Ints(arr1)
	sort.Ints(arr2)
	c := 0
	l, r := 0, 0
	for l < len(arr1) {
		if r == len(arr2) || arr1[l] < arr2[r] {
			c++
			pre = append(pre, arr1[l])
			res = append(res, c)
			ex = append(ex, pre)
			l++
		} else if arr1[l] > arr2[r] {
			r++
		} else {
			res = append(res, c)
			ex = append(ex, pre)
			l++
			r++
		}
	}
	for l < r {
		res = append(res, c)
		ex = append(ex, pre)
		l++
	}

	fmt.Println(ex)
	return res
}
