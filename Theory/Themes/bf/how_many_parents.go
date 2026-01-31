package main

import "fmt"

func LBF1(l, r int, check func(c int, p []int) bool, value ...int) int {
	for l < r {
		m := (l + r) / 2
		if check(m, value) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func checkFunc(m int, p []int) bool {
	n, k := p[0], p[1]
	return (k+m)*3 >= n+m
}

func main() {
	n, k := 9, 2
	fmt.Println(LBF1(0, n, checkFunc, n, k))
}
