package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{171, 189, 170, 179, 181, 183, 174, 175, 175, 175, 187, 186, 183, 181, 169, 172}
	cntPeople, cntGroup := 4, 4
	fmt.Println(comfortHigh(arr, cntPeople, cntGroup))
}

func comfortHigh(arr []int, p, g int) int {
	sort.Ints(arr)
	l, r := 0, arr[len(arr)-1]-arr[0]
	for l < r {
		m := (l + r) / 2
		if approach(m, arr, p, g) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func approach(m int, arr []int, p, g int) bool {
	n := len(arr)
	cnt := 0
	for i := 0; i < n-p+1; {
		if arr[i+p-1]-arr[i] <= m {
			cnt++
			i += p
		} else {
			i++
		}
	}
	return cnt >= g
}
