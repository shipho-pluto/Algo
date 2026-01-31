package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 5, 6, 7, 8}
	fmt.Println(zigzagIterator(arr1, arr2))
}

func zigzagIterator(arr1 []int, arr2 []int) []int {
	n, m := len(arr1), len(arr2)
	res := make([]int, m+n)
	for i, l, r := 0, 0, 0; r < m || l < n; i++ {
		if l < n && r < m {
			if i%2 == 0 {
				res[i] = arr1[l]
				l++
			} else {
				res[i] = arr2[r]
				r++
			}
		} else if l < n {
			res[i] = arr1[l]
			l++
		} else {
			res[i] = arr2[r]
			r++
		}
	}
	return res
}
