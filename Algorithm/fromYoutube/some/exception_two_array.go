package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 6, 6, 8, 91}
	arr2 := []int{1, 5, 6, 10}
	fmt.Println(exception(arr1, arr2))
}

func exception(arr1 []int, arr2 []int) []int {
	var res []int
	n, m := len(arr1), len(arr2)
	for l, r := 0, 0; l < n; {
		if l > 0 && arr1[l] == arr1[l-1] {
			l++
		}

		if r == m || r < m && arr1[l] < arr2[r] {
			res = append(res, arr1[l])
			l++
		} else if r < m && arr1[l] > arr2[r] {
			r++
		} else {
			l, r = l+1, r+1
		}

	}
	return res
}
