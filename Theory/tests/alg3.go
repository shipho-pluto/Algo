package main

import (
	"fmt"
	"sort"
)

func LBF(l, r int, goods []int, w int) int {
	for l < r {
		m := (r + l) / 2
		if goods[m] >= w {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func RBF(l, r int, goods []int, w int) int {
	for l < r {
		m := (r + l + 1) / 2
		if goods[m] < w {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func cntBadDecision(goods, wanted []int) int {
	sort.Slice(goods, func(i, j int) bool {
		return goods[i] < goods[j]
	})
	sum := 0
	n, m := len(goods), len(wanted)
	for i := 0; i < m; i++ {
		sum += min(abs(wanted[i]-goods[LBF(0, n-1, goods, wanted[i])]), abs(wanted[i]-goods[RBF(0, n-1, goods, wanted[i])]))
	}
	return sum
}

func main() {
	goods := []int{1, 8, 3, 5}
	wanted := []int{1, 8, 3, 5}
	fmt.Println(cntBadDecision(goods, wanted))
}
