package main

import "fmt"

func SumIntervals(a1, a2 [][]int) [][]int {
	l, r := 0, 0
	var res [][]int
	for l < len(a1) && r < len(a2) {
		start := max(a1[l][0], a2[r][0])
		end := min(a1[l][1], a2[r][1])

		if end > start {
			res = append(res, []int{start, end})
		}
		if a1[l][1] < a2[r][1] {
			l++
		} else {
			r++
		}
	}

	return res
}

func main() {
	a1, a2 := [][]int{{1, 5}, {14, 20}, {23, 24}}, [][]int{{4, 6}, {10, 14}, {17, 19}}
	fmt.Println(SumIntervals(a1, a2))
}
