package main

import (
	"fmt"
)

func sum(arr []int) int {
	summa := 0
	for i := 0; i < len(arr); i++ {
		summa += arr[i]
	}
	return summa
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxOneElement(arr []int) int {
	p := sum(arr)
	n := len(arr)
	maxLen := 0
	if n == p || p+1 == n {
		return n - 1
	}
	for l := 0; l < n; l++ {
		r, cntN := l, 0
		for r < n {
			if arr[r] == 0 {
				cntN++
			}
			if cntN >= 2 {
				break
			}
			r++
		}
		maxLen = Max(maxLen, r-l-1)
	}
	return maxLen
}

func main() {
	arr := []int{0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1}
	fmt.Println(maxOneElement(arr))
}
