package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{-10, -9, -8},
	}

	fmt.Println(median(arr))
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}

func median(arr [][]int) []int {
	n := len(arr)
	res := make([]int, 0, n*(n-1)/2)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			res = append(res, medianOfTwo(arr[i], arr[j]))
		}
	}

	return res
}

func medianOfTwo(a []int, b []int) int {
	trueM := float64(0)
	for i := range a {
		trueM += float64(a[i])
	}
	for i := range b {
		trueM += float64(b[i])
	}
	trueM = trueM / float64(len(a)+len(b))

	a1, a2 := lbsForM(0, len(a)-1, a, trueM), rbsForM(0, len(a)-1, a, trueM)
	b1, b2 := lbsForM(0, len(b)-1, b, trueM), rbsForM(0, len(b)-1, b, trueM)
	nearA, nearB := 0, 0
	if trueM-float64(a[a2]) < float64(a[a1])-trueM {
		nearA = a[a2]
	} else {
		nearA = a[a1]
	}

	if trueM-float64(b[b2]) < float64(b[b1])-trueM {
		nearB = b[b2]
	} else {
		nearB = b[b1]
	}

	if abs(trueM-float64(nearA)) < abs(trueM-float64(nearB)) {
		return nearA
	}
	return nearB
}

func rbsForM(l int, r int, list []int, t float64) int {
	for l < r {
		m := (l + r + 1) / 2
		if float64(list[m]) > t {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func lbsForM(l int, r int, list []int, t float64) int {
	for l < r {
		m := (l + r) / 2
		if float64(list[m]) < t {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
