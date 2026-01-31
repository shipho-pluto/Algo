package main

import "fmt"

func main() {
	arr := []int{-10, -1, 91, 102, 1229}
	x := 0
	fmt.Println(findNearElement(arr, x))
}

func findNearElement(arr []int, x int) int {
	n := len(arr)
	l := LBF(0, n-1, arr, x)
	r := RBF(0, n-1, arr, x)
	if arr[l]-x < x-arr[r] {
		return arr[l]
	}
	return arr[r]
}

func LBF(l, r int, arr []int, x int) int {
	for l < r {
		m := (l + r) / 2
		if arr[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func RBF(l, r int, arr []int, x int) int {
	for l < r {
		m := (l + r + 1) / 2
		if arr[m] <= x {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
