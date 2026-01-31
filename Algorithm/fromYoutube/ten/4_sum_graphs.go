package main

import "fmt"

func main() {
	arr1 := [][]int{{1, 1}, {3, 2}, {5, 1}}
	arr2 := [][]int{{0, 2}, {4, 4}}
	fmt.Println(sumGraphs(arr1, arr2))
}

func sumGraphs(arr1 [][]int, arr2 [][]int) [][]int {
	var res [][]int
	n, m := len(arr1), len(arr2)
	v1, v2 := 0, 0
	for l, r := 0, 0; l < n || r < m; {
		var point []int
		if l < n && r < m {
			if arr1[l][0] < arr2[r][0] {
				point = []int{arr1[l][0], arr1[l][1] + v2}
				v1 = arr1[l][1]
				l++
			} else if arr1[l][0] > arr2[r][0] {
				point = []int{arr2[r][0], arr2[r][1] + v1}
				v2 = arr2[r][1]
				r++
			} else {
				point = []int{arr1[l][0], arr1[l][1] + arr2[r][1]}
				v1, v2 = arr1[l][1], arr2[r][1]
				l, r = l+1, r+1
			}
		} else if l < n {
			point = []int{arr1[l][0], arr1[l][1] + v2}
			l++
		} else {
			point = []int{arr2[r][0], arr2[r][1] + v1}
			r++
		}
		res = append(res, point)
	}

	return res
}
