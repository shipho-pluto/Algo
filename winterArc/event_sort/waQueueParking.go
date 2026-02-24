package main

import (
	"fmt"
	"sort"
)

type CarQA struct {
	start    int
	end      int
	placeFor int
	placeTo  int
}

func main() {
	arr := []CarQA{{1, 4, 1, 4}, {2, 7, 5, 6}, {1, 9, 7, 10}}
	n := 10
	fmt.Println(placeAtParking(arr, n))
}

func placeAtParking(arr []CarQA, n int) (int, bool) {
	k := len(arr)
	st := make([][]int, 0, 2*k)
	for i := range arr {
		st = append(st, []int{arr[i].start, arr[i].placeTo - arr[i].placeFor})
		st = append(st, []int{arr[i].end, -(arr[i].placeTo - arr[i].placeFor)})
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	res := n
	sum := n
	for i := range st {
		sum -= st[i][1]
		if sum < 0 {
			return n, false
		}
		res = min(res, sum)
	}

	return res, true
}
