package main

import "fmt"

func main() {
	arr := []int{-100, -19, -8, 0, 3, 14, 59, 129}
	fmt.Println(findLenSubarray(arr))
}

func findLenSubarray(arr []int) int {
	n := len(arr)
	return moreZero(0, n-1, arr) - lessZero(0, n-1, arr) + 1
}

func lessZero(l, r int, arr []int) int {
	for l < r {
		m := (l + r) / 2
		if arr[m] >= 0 {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func moreZero(l, r int, arr []int) int {
	for l < r {
		m := (l + r + 1) / 2
		if arr[m] <= 0 {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
