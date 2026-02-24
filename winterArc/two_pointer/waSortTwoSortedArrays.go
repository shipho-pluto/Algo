package main

import "fmt"

func main() {
	arr1, arr2 := []int{1, 4, 9, 10, 11}, []int{-10, -2, 4}
	fmt.Println(sortTwoSorted(arr1, arr2))
}

func sortTwoSorted(arr1 []int, arr2 []int) []int {
	res := make([]int, len(arr1)+len(arr2))

	l, r := 0, 0
	for i := 0; i < len(res); i++ {
		if l < len(arr1) && (r < len(arr2) && arr1[l] <= arr2[r] || r == len(arr2)) {
			res[i] = arr1[l]
			l++
		} else {
			res[i] = arr2[r]
			r++
		}
	}
	return res
}
