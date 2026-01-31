package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1, arr2 := []int{1, 3, 4, 0}, []int{1, 4, 6, 0}
	fmt.Println(eio(arr2, arr1))
	fmt.Println(eioMap(arr2, arr1))
}

func eioMap(arr1 []int, arr2 []int) []int {
	var res []int
	st := make(map[int]bool)
	for i := range arr2 {
		st[arr2[i]] = true
	}

	for i := range arr1 {
		if !st[arr1[i]] {
			res = append(res, arr1[i])
		}
	}
	return res
}

func eio(arr1 []int, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var res []int
	for l, r := 0, 0; l < len(arr1); {
		if l > 0 && arr1[l] == arr1[l-1] {
			l++
		}
		if r == len(arr2) || arr2[r] > arr1[l] {
			res = append(res, arr1[l])
			l++
		} else if arr2[r] < arr1[l] {
			r++
		} else {
			r++
			l++
		}
	}
	return res
}
