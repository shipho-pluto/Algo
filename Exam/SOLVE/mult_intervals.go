package main

import "fmt"

func multiplicationIntervals(in1, in2 [][]int) [][]int {
	n := len(in1)
	var res [][]int
	for l, r := 0, 0; l < n && r < n; {
		start := max(in1[l][0], in2[r][0])
		end := min(in1[l][1], in2[r][1])

		if start < end {
			res = append(res, []int{start, end})
		}

		if in1[l][1] < in2[r][1] {
			l++
		} else {
			r++
		}
	}

	return res
}

func main() {
	a1, a2 := [][]int{{1, 10}, {14, 20}}, [][]int{{2, 5}, {6, 7}}
	fmt.Println(multiplicationIntervals(a1, a2))
}
