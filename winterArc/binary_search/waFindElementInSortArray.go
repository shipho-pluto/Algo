package main

import "fmt"

func main() {
	arr := []int{-10, -1, 0, 91, 102, 1229}
	x := 0
	fmt.Println(findNearElement(arr, x))
}

func findNearElement(arr []int, x int) (int, int) {
	i1, i2 := leftBS(0, len(arr)-1, arr, x), rightBS(0, len(arr)-1, arr, x)
	if arr[i1]-x < x-arr[i2] {
		return arr[i1], i1
	}
	return arr[i2], i2
}

func leftBS(l, r int, arr []int, x int) int {
	for l < r {
		m := (r + l) / 2
		if arr[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func rightBS(l, r int, arr []int, x int) int {
	for l < r {
		m := (r + l + 1) / 2
		if arr[m] > x {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}
