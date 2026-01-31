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

func placeAtParking(arr []CarQA, N int) (int, bool) {
	n := len(arr)
	st := make([][]int, n*2)
	idx := 0
	for i := range n {
		st[idx] = []int{arr[i].start, arr[i].placeTo - arr[i].placeFor + 1, 1}
		idx++
		st[idx] = []int{arr[i].end, arr[i].placeTo - arr[i].placeFor + 1, -1}
		idx++
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][2] < st[j][2]
		}
		return st[i][0] < st[j][0]
	})

	cnt, c := 0, 0
	res := n + 1
	for i := range st {
		cnt += st[i][2] * st[i][1]
		c += st[i][2]
		if cnt == N {
			res = min(res, c)
		}
	}
	if res != n+1 {
		return res, true
	}

	return 0, false
}
