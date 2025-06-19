package main

import "fmt"

func LBF3(l, r int, arr []int, x int) int {
	for l < r {
		m := (l + r) / 2
		if arr[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	if len(arr)-1 == l && arr[l] != x {
		return l + 1
	}
	return l
}

func main() {
	arr := []int{-10, -3, 0, 4, 8, 10, 19, 24, 98}
	x := 99
	fmt.Println(LBF3(0, len(arr)-1, arr, x))
}
