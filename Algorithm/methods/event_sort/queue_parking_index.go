package main

import (
	"fmt"
	"sort"
)

type CarQueue struct {
	start, end           int
	placeStart, placeEnd int
}

func main() {
	arr := []CarQueue{
		{2, 7, 5, 6},
		{1, 9, 7, 10},
		{6, 8, 1, 2},
		{6, 8, 2, 3},
	}
	n := 10
	fmt.Println(indexMinCar(arr, n))
}

func indexMinCar(arr []CarQueue, N int) []int {
	n := len(arr)
	st := make([][]int, 2*n)
	idx := 0
	for i := range n {
		st[idx] = []int{arr[i].start, arr[i].placeEnd - arr[i].placeStart + 1, 1, i}
		idx++
		st[idx] = []int{arr[i].end, arr[i].placeEnd - arr[i].placeStart + 1, -1, i}
		idx++
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][2] < st[j][2]
		}
		return st[i][0] < st[j][0]
	})

	var res []int
	index := make(map[int]bool, n)
	minCnt := n + 1
	cnt, c := 0, 0
	for i := range st {
		c += st[i][2]
		if st[i][2] > 0 {
			index[st[i][3]] = true
		} else {
			delete(index, st[i][3])
		}
		cnt += st[i][2] * st[i][1]
		if cnt == N {
			if c < minCnt {
				minCnt = c
				res = []int{}
				for k := range index {
					res = append(res, k)
				}
			}
		}
	}
	return res
}
