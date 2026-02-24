package main

import (
	"fmt"
)

func main() {
	arr := []int{-1, -2, -4, -2, 0, 1, 1, 4, 5, 7}
	fmt.Println(minMul(arr))

	arr = []int{5, 6, 12, 13, 20, 21}
	k, i := 4, 3
	fmt.Println(findCloset(arr, k, i))

	a, b := [][]int{{1, 2}, {5, 3}}, [][]int{{2, 6}, {3, 8}, {9, 3}}
	fmt.Println(sumGraph(a, b))
}

func sumGraph(a [][]int, b [][]int) any {
	v1, v2 := 0, 0
	var res [][]int
	for l, r := 0, 0; l < len(a) || r < len(b); {
		var point []int
		if l < len(a) && r < len(b) {
			if a[l][0] < b[r][0] {
				point = []int{a[l][0], a[l][1] + v2}
				v1 = a[l][1]
				l++
			} else if a[l][0] > b[r][0] {
				point = []int{b[r][0], b[r][1] + v1}
				v2 = b[r][1]
				r++
			} else {
				point = []int{b[r][0], b[r][1] + a[l][1]}
				v1, v2 = a[l][1], b[r][1]
				l, r = l+1, r+1
			}
		} else if l < len(a) {
			point = []int{a[l][0], a[l][1] + v2}
			l++
		} else {
			point = []int{b[r][0], b[r][1] + v1}
			r++
		}

		res = append(res, point)
	}

	return res
}

func findCloset(arr []int, k int, i int) any {
	res := make([]int, 0, k)
	for l, r := i, i+1; (l > 0 || r < len(arr)) && k > 0; {
		if l > 0 && r < len(arr) && arr[i]-arr[l] <= arr[r]-arr[i] || r == len(arr) {
			res = append(res, arr[l])
			l--
			k--
		} else if r < len(arr) {
			res = append(res, arr[r])
			r++
			k--
		}
	}
	return res
}

func minMul(arr []int) int {
	max1, max2 := -10000000000001, -10000000000000
	min1, min2 := 10000000000001, 10000000000000
	isLess, isMore := false, false
	for i := range arr {
		if arr[i] > max1 {
			max2 = max1
			max1 = arr[i]
		} else if arr[i] > max2 {
			max2 = arr[i]
		}

		if arr[i] < min1 {
			min2 = min1
			min1 = arr[i]
		} else if arr[i] < min2 {
			min2 = arr[i]
		}

		if arr[i] > 0 && !isMore {
			isMore = true
		} else if arr[i] < 0 && !isLess {
			isLess = true
		}
	}

	if isMore && isLess {
		return max1 * min1
	} else if isMore {
		return min1 * min2
	}
	return max2 * max1
}
