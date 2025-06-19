package main

import "fmt"

func MergeTwoArrayInOne(a1, a2 []int) []int {
	n, m := len(a1), len(a2)
	res := make([]int, n+m)
	l, r, i := 0, 0, 0
	for ; l < n && r < m; i++ {
		if a1[l] <= a2[r] {
			res[i] = a1[l]
			l++
		} else {
			res[i] = a2[r]
			r++
		}
	}

	for ; r < m; func() { r++; i++ }() {
		res[i] = a2[r]
	}
	for ; l < n; func() { l++; i++ }() {
		res[i] = a1[l]
	}
	return res
}

func main() {
	arr1, arr2 := []int{3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 7, 8, 12}, []int{-4, -1, 0, 1, 2, 4, 5, 8, 10, 11, 16, 18, 21}
	fmt.Println(MergeTwoArrayInOne(arr1, arr2))
}
