package main

import "fmt"

func main() {
	arr1, arr2 := []int{1, 3, 5}, []int{2, 4}
	fmt.Println(mta(arr1, arr2))
}

func mta(arr1 []int, arr2 []int) []int {
	n, m := len(arr1), len(arr2)
	res := make([]int, n+m)
	for i, l, r := 0, 0, 0; l < n || r < m; i++ {
		if l < n && r < m {
			if arr1[l] < arr2[r] {
				res[i] = arr1[l]
				l++
			} else {
				res[i] = arr2[r]
				r++
			}
		} else if r < m {
			res[i] = arr2[r]
			r++
		} else {
			res[i] = arr1[l]
			l++
		}
	}

	return res
}
