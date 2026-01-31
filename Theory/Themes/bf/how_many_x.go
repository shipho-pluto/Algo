package main

import "fmt"

func help(l, r int, arr []int, x int) int {
	return LBF42(l, r, arr, x) - LBF41(l, r, arr, x)
}

func LBF41(l, r int, arr []int, x int) int {
	for l < r {
		m := (l + r) / 2
		if arr[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	if len(arr) == l+1 {
		return l + 1
	}
	return l
}

func LBF42(l, r int, arr []int, x int) int {
	for l < r {
		m := (l + r) / 2
		if arr[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	if len(arr) == l+1 {
		return l + 1
	}
	return l
}

func main() {
	arr := []int{-10, -10, -3, 0, 4, 8, 10, 19, 24, 98, 100, 100, 101, 101}
	x := 0
	fmt.Println(help(0, len(arr)-1, arr, x))
}
