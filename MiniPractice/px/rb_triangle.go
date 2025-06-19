package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x, y int
}

func distance(x1, y1, x2, y2 int) int {
	return int(math.Pow(math.Abs(float64(x1)-float64(x2)), 2) + math.Pow(math.Abs(float64(y1)-float64(y2)), 2))
}

func RBTriangle(arr []Point) int {
	var pref [][]int
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			pref = append(pref, []int{distance(arr[i].x, arr[i].y, arr[j].x, arr[j].y), int(math.Abs(float64(arr[i].y-arr[j].y)) / float64(arr[i].x-arr[j].x))})
		}
	}

	sort.Slice(pref, func(i, j int) bool {
		return pref[i][0] < pref[i][0]
	})

	r, res, cnt := 1, 0, 0
	for l := 0; l < n-1; {
		for r < n && pref[r][0]-pref[l][0] == 0 {
			if pref[r][1] == pref[l][1] {
				cnt++
			}
			r++
		}
		res += r - l - cnt
		l = r
		r++
	}
	return res
}

func main() {
	arr := []Point{Point{0, 0}, Point{-2, -2}, Point{-2, 2}, Point{2, 2}, Point{1, 0}, {0, -1}}
	fmt.Println(RBTriangle(arr))
}
