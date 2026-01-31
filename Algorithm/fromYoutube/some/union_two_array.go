package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 6, 6, 8, 10}
	arr2 := []int{1, 5, 6, 10}
	fmt.Println(union(arr1, arr2))
}

func union(arr1 []int, arr2 []int) []int {
	var res []int
	n, m := len(arr1), len(arr2)
	for l, r := 0, 0; l < n && r < m; {
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
