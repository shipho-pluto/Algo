package main

import "fmt"

func merge(n1, n2 []int, m int) (int, int) {
	less := lbf(0, len(n1), n1, m) + lbf(0, len(n2), n2, m)
	return less, len(n1)*2 - less
}

func lbf(l, r int, n1 []int, val int) int {
	for l < r {
		m := (r + l) / 2
		if n1[m] >= val {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func LBFn(l, r int, check func(n1, n2 []int, m int) (int, int), n1, n2 []int) int {
	for l < r {
		m := (r + l) / 2
		less, great := check(n1, n2, m)
		if less <= len(n1) && great <= len(n2)-1 {
			return m
		} else if less > len(n1) {
			r = m - 1
		} else if great > len(n2)-1 {
			l = m + 1
		}
	}
	return l
}

func median(n1, n2 []int) int {
	return LBFn(min(n1[0], n2[0]), max(n1[len(n1)-1], n2[len(n2)-1]), merge, n1, n2)
}

func main() {
	n := [][]int{
		{1, 2, 3},
		{1, 3, 3},
		{-10, -9, -8},
	}

	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			fmt.Println(median(n[i], n[j]))
		}
	}
}
