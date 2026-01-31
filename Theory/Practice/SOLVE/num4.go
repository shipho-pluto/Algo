package main

import "fmt"

func SumTwoArr(arr1, arr2 [][]int) [][]int {
	l, r := 0, 0
	n, m := len(arr1), len(arr2)
	v1, v2 := 0, 0
	var res [][]int
	for l < n || r < m {
		point := make([]int, 2)
		if l < n && r < m {
			if arr1[l][0] < arr2[r][0] {
				point[0] = arr1[l][0]
				point[1] = v2 + arr1[l][1]
				v1 = arr1[l][1]
				l++
			} else if arr1[l][0] > arr2[r][0] {
				point[0] = arr2[r][0]
				point[1] = v1 + arr2[r][1]
				v2 = arr2[r][1]
				r++
			} else {
				point[0] = arr2[r][0]
				point[1] = arr1[l][1] + arr2[r][1]
				v2 = arr2[r][1]
				v1 = arr1[l][1]
				r++
				l++
			}
		} else if l < n {
			point[0] = arr1[l][0]
			point[1] = arr1[l][1] + v2
			l++
		} else if r < m {
			point[0] = arr2[r][0]
			point[1] = arr2[r][1] + v1
			r++
		}
		res = append(res, point)
	}

	return res
}

func main() {
	arr1 := [][]int{{1, 2}, {4, 3}, {5, 5}, {7, 9}}
	arr2 := [][]int{{0, 1}, {4, 5}, {7, 2}, {11, 3}}
	fmt.Println(SumTwoArr(arr1, arr2))
}
