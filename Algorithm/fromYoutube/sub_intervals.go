package main

import "fmt"

func main() {
	a1, a2 := [][]int{{1, 5}, {14, 20}, {23, 24}}, [][]int{{4, 6}, {10, 14}, {17, 19}}
	fmt.Println(sumIntervals(a1, a2))
}

func sumIntervals(a1 [][]int, a2 [][]int) [][]int {
	var res [][]int
	n, m := len(a1), len(a2)
	for l, r := 0, 0; l < n && r < m; {
		end := min(a1[l][1], a2[r][1])
		start := max(a1[l][0], a2[r][0])

		if start < end {
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
