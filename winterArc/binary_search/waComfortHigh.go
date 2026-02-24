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

func comfortHigh(arr []int, p int, g int) int {
	sort.Ints(arr)
	return lbsForComfortHigh(0, arr[len(arr)-1]-arr[0], arr, p, g)
}

func lbsForComfortHigh(l int, r int, arr []int, p int, g int) int {
	for l < r {
		m := (l + r) / 2
		if checkForComfortHigh(m, p, g, arr) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func checkForComfortHigh(m int, p int, g int, arr []int) bool {
	cnt := 0
	for i := p - 1; i < len(arr); {
		if arr[i]-arr[i+1-p] <= m {
			cnt++
			i += p
		} else {
			i++
		}
	}

	return cnt < g
}
