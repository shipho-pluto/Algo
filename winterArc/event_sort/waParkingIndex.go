package main

import (
	"fmt"
	"sort"
)

type CarQB struct {
	start, end           int
	placeStart, placeEnd int
}

func main() {
	arr := []CarQB{{2, 7, 5, 6}, {1, 9, 7, 10},
		{6, 8, 1, 2}, {6, 8, 2, 3}}
	n := 10
	fmt.Println(indexMinCar(arr, n))
}

func indexMinCar(arr []CarQB, n int) any {
	k := len(arr)
	st := make([][]int, 0, 2*k)
	for i := range arr {
		st = append(st, []int{arr[i].start, arr[i].placeEnd - arr[i].placeStart, i})
		st = append(st, []int{arr[i].end, -(arr[i].placeEnd - arr[i].placeStart), i})
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	var res []int
	index := make(map[int]int)
	minEl := n
	minInd := 0
	for i := range st {
		if st[i][1] > 0 {
			index[st[i][2]] = st[i][1]
			if minEl > st[i][1] {
				minEl = st[i][1]
				minInd = st[i][2]
				res = append(res, minInd)
			}
		} else {
			delete(index, st[i][2])
			if minInd == st[i][2] && i != len(st)-1 {
				curEl := n
				curInd := -1
				for k, v := range index {
					if curEl > v {
						curEl = v
						curInd = k
					}
				}
				minInd = curInd
				minEl = curEl
				res = append(res, minInd)
			}
		}
	}

	return res
}
