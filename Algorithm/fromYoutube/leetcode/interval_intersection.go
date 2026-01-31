package main

import "fmt"

func main() {
	arr1, arr2 := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	fmt.Println(intervalIntersection(arr1, arr2))
}
func intervalIntersection(arr1 [][]int, arr2 [][]int) [][]int {
	n, m := len(arr1), len(arr2)
	var res [][]int
	for l, r := 0, 0; l < n && r < m; {
		start := max(arr1[l][0], arr2[r][0])
		end := min(arr1[l][1], arr2[r][1])

		if start <= end {
			res = append(res, []int{start, end})
		}
		if arr1[l][1] < arr2[r][1] {
			l++
		} else {
			r++
		}
	}
	return res
}
