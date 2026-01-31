package main

import "fmt"

func LBF2(l, r int, check func(c int, p []int) bool, param ...int) int {
	for l < r {
		m := (r + l) / 2
		if check(m, param) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func check(m int, param []int) bool {
	n, k := param[0], param[1]
	return (2*k+m-1)*m/2 >= n
}

func main() {
	n, k := 49, 4
	fmt.Println(LBF2(0, n, check, n, k))
}
