package main

func LBF(l, r int, check func(c int, v []int) bool, value ...int) int {
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

func RBF(l, r int, check func(c int, v []int) bool, value ...int) int {
	for l < r {
		m := (r + l + 1) / 2
		if check(m, value) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
