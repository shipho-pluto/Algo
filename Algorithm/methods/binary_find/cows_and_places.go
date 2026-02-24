package main

import (
	"fmt"
	"sort"
)

func main() {
	n, cows := []int{2, 5, 7, 11, 15, 20}, 3
	fmt.Println(howManyPlaces(1, n[len(n)-1], n, cows))
}

func howManyPlaces(l int, r int, n []int, cows int) any {
	sort.Ints(n)
	for l < r {
		m := (l + r + 1) / 2
		if approved(n, m, cows) {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func approved(arr []int, m, c int) bool {
	cnt := 1
	for l, r := 0, 1; r < len(arr); r++ {
		if arr[r]-arr[l] >= m {
			cnt++
			l = r
		}
	}
	return cnt < c
}
