package main

import "fmt"

func main() {
	cntPeople, cntParents := 9, 2
	fmt.Println(ePTP(0, cntPeople, check, cntPeople, cntParents))
}

func ePTP(l int, r int, c func(m int, param []int) bool, val ...int) int {
	for l < r {
		m := (r + l) / 2
		if c(m, val) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func check(m int, param []int) bool {
	n, k := param[0], param[1]
	return (k + m) > (n+m)/3.0
}
