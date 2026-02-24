package main

import "fmt"

func main() {
	arr1 := [][]int{{0, 1}, {1, 3}, {3, 2}, {7, 7}}
	arr2 := [][]int{{1, 2}, {6, 5}, {8, 4}, {9, 5}, {10, 6}}

	fmt.Println(sumGraphic(arr1, arr2))
}

func sumGraphic(arr1 [][]int, arr2 [][]int) [][]int {
	var res [][]int
	v1, v2 := 0, 0
	for l, r := 0, 0; l < len(arr1) || r < len(arr2); {
		var point []int
		if l < len(arr1) && r < len(arr2) {
			if arr1[l][0] < arr2[r][0] {
				point = []int{arr1[l][0], arr1[l][1] + v2}
				v1 = arr1[l][1]
				l++
			} else if arr1[l][0] > arr2[r][0] {
				point = []int{arr2[r][0], arr2[r][1] + v1}
				v2 = arr2[r][1]
				r++
			} else {
				point = []int{arr2[r][0], arr2[r][1] + arr1[l][1]}
				v2 = arr2[r][1]
				v1 = arr1[l][1]
				l++
				r++
			}
		} else if l < len(arr1) {
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
