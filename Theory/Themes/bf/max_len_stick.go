package main

import "fmt"

func RBF1(l, r int, check func(c int, p []int) bool, param ...int) int {
	for l < r {
		m := (r + l + 1) / 2
		if check(m, param) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func maxLenStick(m int, param []int) bool {
	w, h, n := param[0], param[1], param[2]
	return (w/m)*(h/m) >= n
}

func main() {
	w, h, n := 10, 20, 4
	fmt.Println(RBF1(0, w, maxLenStick, w, h, n))
}
