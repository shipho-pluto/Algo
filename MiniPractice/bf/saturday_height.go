package main

import (
	"fmt"
	"sort"
)

func ComfortHeight(middle int, arr []int, p []int) bool {
	r, c := p[0], p[1]
	cnt := 0
	for i := 0; i < len(arr)-c+1; {
		if arr[i+c-1]-arr[i] <= middle {
			cnt++
			i += c
		} else {
			i++
		}
	}
	return cnt >= r
}

func LBF4(l, r int, check func(c int, arr, m []int) bool, arr []int, val ...int) int {
	for l < r {
		m := (l + r) / 2
		if check(m, arr, val) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func Saturday(n []int, r, c int) int {
	sort.Ints(n)
	return LBF4(0, n[len(n)-1]-n[0], ComfortHeight, n, r, c)
}

func main() {
	n := []int{171, 189, 170, 179, 181, 183, 174, 175, 175, 175, 187, 186, 183, 181, 169, 172}
	r, c := 4, 4
	fmt.Println(Saturday(n, r, c))
}
