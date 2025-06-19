package main

import "fmt"

type Point3 struct {
	t, v int
}

func SumGraphs(g1 []Point3, g2 []Point3) []Point3 {
	var res []Point3
	l, r := 0, 0
	n, m := len(g1), len(g2)
	val1, val2 := 0, 0
	for l < n || r < m {
		pt := Point3{}
		if l < n && r < m {
			if g1[l].t < g2[r].t {
				pt = Point3{g1[l].t, g1[l].v + val2}
				val1 = g1[l].v
				l++
			} else {
				pt = Point3{g2[r].t, g2[r].v + val1}
				val2 = g2[r].v
				r++
			}

		} else if l < n {
			pt = Point3{g1[l].t, g1[l].v + val2}
			l++
		} else if r < m {
			pt = Point3{g2[r].t, g2[r].v + val1}
			r++
		}
		if len(res) > 0 && res[len(res)-1].v != pt.v {
			res = append(res, pt)
		} else if len(res) == 0 {
			res = append(res, pt)
		}

	}

	return res
}

func main() {
	arr1 := []Point3{Point3{0, 2}, Point3{2, 3}, Point3{6, 4}}
	arr2 := []Point3{Point3{1, 0}, Point3{5, 1}, Point3{7, 2}, Point3{9, 3}}
	fmt.Println(SumGraphs(arr1, arr2))
}
