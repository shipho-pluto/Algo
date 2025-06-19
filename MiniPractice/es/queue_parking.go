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

func PlaceAtParking(arr []CarQA, N int) (int, bool) {
	var events [][]int
	n := len(arr)
	for i := 0; i < n; i++ {
		st, ed, pf, pt := arr[i].start, arr[i].end, arr[i].placeFor, arr[i].placeTo
		events = append(events, []int{st, 1, pt - pf + 1})
		events = append(events, []int{ed, -1, pt - pf + 1})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			if events[i][1] == events[j][1] {
				return events[i][2] < events[j][2]
			}
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})

	cnt, res := 0, n+1
	c := 0
	for _, v := range events {
		cnt += v[1] * v[2]
		c += v[1]
		if cnt == N {
			res = min(res, c)
		}
	}

	if res != n+1 {
		return res, true
	}
	return 0, false
}

func main() {
	arr := []CarQA{{1, 4, 1, 4}, {2, 7, 5, 6}, {1, 9, 7, 10}}
	n := 10
	fmt.Println(PlaceAtParking(arr, n))
}
