package main

import "fmt"

func main() {
	arr1 := [][]int{{0, 1}, {1, 3}, {3, 2}, {7, 7}}
	arr2 := [][]int{{1, 2}, {6, 5}, {8, 4}, {9, 5}, {10, 6}}
	fmt.Println(sumGraphs(arr1, arr2))
}

func sumGraphs(arr1 [][]int, arr2 [][]int) [][]int {
	m, n := len(arr1), len(arr2)
	var res [][]int
	v1, v2 := 0, 0
	for l, r := 0, 0; r < n || l < m; {
		newPoint := make([]int, 2)
		if l < m && r < n {
			if arr1[l][0] < arr2[r][0] {
				newPoint = []int{arr1[l][0], arr1[l][1] + v2}
				v1 = arr1[l][1]
				l++
			} else if arr1[l][0] > arr2[r][0] {
				newPoint = []int{arr2[r][0], arr2[r][1] + v1}
				v2 = arr2[r][1]
				r++
			} else {
				newPoint = []int{arr1[l][0], arr1[l][1] + arr2[r][1]}
				v1, v2 = arr1[l][1], arr2[r][1]
				l, r = l+1, r+1
			}
		} else if l < m {
			newPoint = []int{arr1[l][0], arr1[l][1] + v2}
			l++
		} else {
			newPoint = []int{arr2[r][0], arr2[r][1] + v1}
			r++
		}

		res = append(res, newPoint)
	}
	return res
}
