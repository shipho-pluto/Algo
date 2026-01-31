package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{1, 3, 3},
		{-10, -9, -8},
	}
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fmt.Println(medianTwoArray(arr[i], arr[j]))
		}
	}
}

func medianTwoArray(arr1, arr2 []int) int {
	return fbM2A(min(arr1[0], arr2[0]), max(arr1[len(arr1)-1], arr2[len(arr2)-1]), arr1, arr2)
}

func fbM2A(l, r int, arr1, arr2 []int) int {
	for l < r {
		m := (r + l) / 2
		if median(m, arr1, arr2) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func median(mid int, arr1, arr2 []int) bool {
	n, m := len(arr1), len(arr2)
	sum := 0
	for l := range n {
		sum += arr1[l]
	}
	for r := range m {
		sum += arr2[r]
	}
	return (mid) >= sum/(n+m)
}
