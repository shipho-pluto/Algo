package main

import (
	"fmt"
	"sort"
)

func main() {
	cntPeople := 9
	cntParents := 2
	k := 0.3
	fmt.Println(cntParentsMoreK(cntPeople, cntParents, k))

	arr := []int{-10, -1, 91, 102, 1229}
	x := 0
	fmt.Println(findNearElement(arr, x))

	ar := [][]int{{0, 3}, {5, 10}, {2, 4}, {6, 8}, {11, 12}, {16, 21}, {12, 19}, {20, 23}}
	fmt.Println(lenInter(ar))
}

func lenInter(arr [][]int) int {
	st := make([][]int, 0, len(arr)*2)
	for i := range arr {
		st = append(st, []int{arr[i][0], 1})
		st = append(st, []int{arr[i][1], -1})
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})
	s := 0
	res := 0
	l := 0
	for r := range st {
		s += st[r][1]
		if s == 0 {
			res += st[r][0] - st[l][0]
			l = r + 1
		}
	}

	return res
}

func findNearElement(arr []int, x int) any {
	i1, i2 := lbf(0, len(arr)-1, arr, x), rbf(0, len(arr)-1, arr, x)
	if arr[i1]-x < x-arr[i2] {
		return i1
	}
	return i2
}

func rbf(l, r int, arr []int, x int) int {
	for l < r {
		m := (l + r + 1) / 2
		if arr[m] > x {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func lbf(l, r int, arr []int, x int) int {
	for l < r {
		m := (l + r) / 2
		if arr[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func cntParentsMoreK(n int, p int, k float64) any {
	return lbs(0, n, n, p, k)
}

func lbs(l int, r int, n int, p int, k float64) any {
	for l < r {
		m := (l + r) / 2
		if float64(m+p) < k*float64(n+m) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
