package main

import (
	"fmt"
	"sort"
)

type CarQA2 struct {
	start    int
	end      int
	placeFor int
	placeTo  int
}

func MinAutoParking(arr []CarQA2, N int) []int {
	n := len(arr)
	m := make([][]int, 2*n)
	for i, in := 0, 0; i < n; func() { i++; in += 2 }() {
		m[in] = []int{arr[i].start, 1, arr[i].placeTo - arr[i].placeFor + 1, i}
		m[in+1] = []int{arr[i].end, -1, arr[i].placeTo - arr[i].placeFor + 1, i}
	}

	sort.Slice(m, func(i, j int) bool {
		if m[i][0] == m[j][0] {
			if m[i][1] == m[j][1] {
				return m[i][2] < m[j][2]
			}
			return m[i][1] < m[j][1]
		}
		return m[i][0] < m[j][0]
	})

	var res []int
	c := make(map[int]int)
	lc := 0
	cnt, mn := 0, n+1
	for _, v := range m {
		cnt += v[1] * v[2]
		if v[1] == 1 {
			c[v[3]]++
			lc++
		} else {
			c[v[3]] = 0
			lc--
		}
		if cnt == N {
			if lc < mn {
				mn = lc
				res = []int{}
				for i := range c {
					if c[i] > 0 {
						res = append(res, i)
					}
				}
			}
		}
	}

	return res
}

func main() {
	arr := []CarQA2{{2, 7, 5, 6}, {1, 9, 7, 10}, {6, 8, 1, 2}, {6, 8, 2, 3}}
	n := 10
	fmt.Println(MinAutoParking(arr, n))
}
