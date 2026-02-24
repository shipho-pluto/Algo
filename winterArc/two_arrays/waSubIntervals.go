package main

import "fmt"

func main() {
	a1, a2 := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	fmt.Println(subIntervals(a1, a2))
}

func subIntervals(a1 [][]int, a2 [][]int) any {
	var res [][]int
	for l, r := 0, 0; l < len(a1) && r < len(a2); {
		mxStart := max(a1[l][0], a2[r][0])
		mnEnd := min(a1[l][1], a2[r][1])

		if mxStart < mnEnd {
			res = append(res, []int{mxStart, mnEnd})
		} else if mxStart == mnEnd {
			res = append(res, []int{mxStart})
		}

		if a1[l][1] < a2[r][1] {
			l++
		} else {
			r++
		}
	}

	return res
}
