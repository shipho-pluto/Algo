package main

import (
	"fmt"
)

func abs1(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func AddFirst(element int, s []int) []int {
	res := make([]int, len(s)+1)
	copy(res[1:], s)
	res[0] = element
	return res
}

func main() {
	a := []int{-20, -5, -3, -2, 3, 6, 9, 11}

	aLeft := a[0]
	il, ir := 0, len(a)-1
	aRight := a[len(a)-1]
	var newA []int

	for aLeft <= aRight {
		if abs1(aLeft) > abs1(aRight) {
			newA = AddFirst(aLeft*aLeft, newA)
			il += 1
			aLeft = a[il]
		} else {
			newA = AddFirst(aRight*aRight, newA)
			ir -= 1
			aRight = a[ir]
		}
	}
	fmt.Println(newA)
}
