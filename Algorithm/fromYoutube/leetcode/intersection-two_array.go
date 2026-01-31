package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 2, 1}
	arr2 := []int{2, 2}
	fmt.Println(intersection(arr1, arr2))
	fmt.Println(intersect(arr1, arr2))
}

func intersection(arr1 []int, arr2 []int) []int {
	mask := make(map[int]bool)
	var res []int
	for i := range arr1 {
		mask[arr1[i]] = true
	}
	for i := range arr2 {
		if mask[arr2[i]] {
			res = append(res, arr2[i])
			delete(mask, arr2[i])
		}
	}
	return res
}

func intersect(arr1 []int, arr2 []int) []int {
	mask := make(map[int]int)
	var res []int
	for i := range arr1 {
		mask[arr1[i]]++
	}
	for i := range arr2 {
		if val := mask[arr2[i]]; val > 0 {
			res = append(res, arr2[i])
			mask[arr2[i]]--
		}
	}
	return res
}
