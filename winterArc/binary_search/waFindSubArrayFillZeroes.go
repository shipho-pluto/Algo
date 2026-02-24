package main

import "fmt"

func main() {
	arr := []int{-100, -19, -8, 0, 0, 0, 14, 59, 129}
	fmt.Println(findLenSubarray(arr))
}

func findLenSubarray(arr []int) any {
	return LBS(0, len(arr)-1, arr) - RBS(0, len(arr)-1, arr) + 1
}

func RBS(l int, r int, arr []int) int {
	for l < r {
		m := (r + l) / 2
		if arr[m] < 0 {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func LBS(l int, r int, arr []int) int {
	for l < r {
		m := (r + l + 1) / 2
		if arr[m] > 0 {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}
