package main

import "fmt"

type Point6 struct {
	time  int
	value int
}

func SumTwoGraphs(a1, a2 []Point6) []Point6 {
	var res []Point6
	l, r := 0, 0
	n1, n2 := len(a1), len(a2)
	v1, v2 := 0, 0
	for l < n1 || r < n2 {
		if l < n1 && r < n2 {
			if a1[l].time == a2[r].time {
				v1 = a1[l].value
				v2 = a2[r].value
				res = append(res, Point6{a1[l].time, v1 + v2})
				l++
				r++
			} else if a1[l].time < a2[r].time {
				v1 = a1[l].value
				res = append(res, Point6{a1[l].time, v1 + v2})
				l++
			} else {
				v2 = a2[r].value
				res = append(res, Point6{a2[r].time, v1 + v2})
				r++
			}
		} else if l < n1 {
			res = append(res, Point6{a1[l].time, a1[l].value + v2})
			l++
		} else if r < n2 {
			res = append(res, Point6{a2[r].time, a2[r].value + v1})
			r++
		}
	}
	return res
}

func main() {
	a1 := []Point6{{1, 10}, {5, 20}}
	a2 := []Point6{{2, 100}, {3, 200}, {4, 300}, {6, 400}}
	fmt.Println(SumTwoGraphs(a1, a2))
}
