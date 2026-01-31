package main

import "fmt"

type Point4 struct {
	x, y int
}

func IsLineSeparator(arr []Point4) bool {
	n := len(arr)
	st := map[Point4]int{}
	maxX, minX := arr[0].x, arr[0].x
	for i := 0; i < n; i++ {
		if arr[i].x > maxX {
			maxX = arr[i].x
		}
		if arr[i].x < minX {
			minX = arr[i].x
		}
		st[arr[i]]++
	}

	lineX := maxX - minX

	for pt := range st {
		ptTo := Point4{lineX - pt.x, pt.y}
		if st[pt] != st[ptTo] {
			return false
		}
	}

	return true
}

func main() {
	arr := []Point4{{0, 0}, {0, 1}, {1, 1}, {2, 3}, {2, 2}, {3, 1}, {4, 1}, {4, 0}}
	fmt.Println(IsLineSeparator(arr))
}
