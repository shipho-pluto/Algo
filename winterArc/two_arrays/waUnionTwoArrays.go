package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 6, 6, 8, 10}
	arr2 := []int{1, 5, 6, 10}
	fmt.Println(unionArrays(arr1, arr2))
	fmt.Println(unionArraysLR(arr1, arr2))
}

func unionArrays(arr1 []int, arr2 []int) []int {
	dict := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		dict[arr1[i]]++
	}

	var res []int

	for i := range arr2 {
		if _, ok := dict[arr2[i]]; ok {
			dict[arr2[i]]--
			res = append(res, arr2[i])
			if dict[arr2[i]] == 0 {
				delete(dict, arr2[i])
			}
		}
	}
	return res
}

func unionArraysLR(arr1 []int, arr2 []int) []int {
	var res []int
	for l, r := 0, 0; l < len(arr1) && r < len(arr2); {
		if arr2[r] < arr1[l] {
			r++
		} else if arr2[r] > arr1[l] {
			l++
		} else {
			res = append(res, arr2[r])
			l, r = l+1, r+1
		}
	}

	return res
}
