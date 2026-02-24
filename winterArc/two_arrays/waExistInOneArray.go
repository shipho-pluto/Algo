package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1, arr2 := []int{1, 3, 4, 0, 0}, []int{1, 4, 6, 0}
	fmt.Println(existInOne(arr1, arr2))
	fmt.Println(existInOneCount(arr1, arr2))
	fmt.Println(existInOneLR(arr1, arr2))
	fmt.Println(existInOneLRCount(arr1, arr2))
}

func existInOne(arr1 []int, arr2 []int) []int {
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

func existInOneCount(arr1 []int, arr2 []int) []int {
	st := make(map[int]int)
	var res []int
	for i := range arr2 {
		st[arr2[i]]++
	}
	for i := range arr1 {
		if st[arr1[i]] > 0 {
			st[arr1[i]]--
		} else {
			res = append(res, arr1[i])
		}
	}

	return res
}

func existInOneLR(arr1 []int, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var res []int
	for l, r := 0, 0; l < len(arr1); {
		if r < len(arr2) {
			if arr1[l] < arr2[r] {
				res = append(res, arr1[l])
				l++
			} else if arr1[l] > arr2[r] {
				r++
			} else {
				l++
			}
		} else {
			res = append(res, arr1[l])
			l++
		}
	}

	return res
}

func existInOneLRCount(arr1 []int, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var res []int
	for l, r := 0, 0; l < len(arr1); {
		if r < len(arr2) {
			if arr1[l] < arr2[r] {
				res = append(res, arr1[l])
				l++
			} else if arr1[l] > arr2[r] {
				r++
			} else {
				l, r = l+1, r+1
			}
		} else {
			res = append(res, arr1[l])
			l++
		}
	}

	return res
}
